document
  .getElementById("new-task-form")
  .addEventListener("submit", async function (e) {
    e.preventDefault(); // Prevent page reload

    const form = e.target;
    const input_data = document.getElementById("new-task-input");
    const formData = new FormData();
    formData.append("description", input_data.value);

    const response = await fetch("/todo", {
      method: "POST",
      body: formData,
    });

    if (response.ok) {
      form.reset();
      fetchItems();
    } else {
      alert("Unsuccesful");
    }
  });

async function fetchItems() {
  const response = await fetch("/todo");
  const todos = await response.json();
  const listElement = document.querySelector("#items-list");

  var list = "";
  todos.forEach((todo) => {
    const htmlItem = `
<li data-status="${todo["status"]}" id="${todo["id"]}" class="task">
  <div class="task-div">
      <span class="${todo["status"] === true ? "strikethrough" : ""}" onclick="toggleStatus(this.parentElement.parentElement);">${todo["description"]}</span>
    <button class="task-list-button" onclick="editItemForm(this.parentElement);">
      Edit
    </button>
    <button class="task-list-button" onclick="deleteItem(this.parentElement);">Delete</button>
  </div>
</li>
    `;
    list = list + htmlItem;
  });

  listElement.innerHTML = list;
}

function editFormTemplate(itemDiv) {
  const form = document.createElement("form");
  const item = itemDiv.parentElement;
  const itemId = item.id;
  const status = item.dataset.status;
  const currentText = itemDiv.querySelector("span").textContent.trim();
  form.classList.add("edit-form");
  form.onsubmit = (e) => {
    e.preventDefault();
    editItem(e.target);
  };
  form.innerHTML = `
  <input id="hidden-id" type="hidden" name="id" value="${itemId}" />
  <input id="hidden-status" type="hidden" name="status" value="${status}" />
  <input id="description" type="text" name="description" class="edit-form-input" value="${currentText}" />
      <button type="submit" class="task-list-button">
        Save
      </button>
      <button type="button" class="task-list-button" onclick="cancelEdit(this.parentElement.parentElement)">
        Cancel
      </button>
  `;
  return form;
}

function editItemForm(itemDiv) {
  const form = editFormTemplate(itemDiv);
  itemDiv.parentElement.appendChild(form);
  itemDiv.classList.add("hidden");
}

async function editItem(form) {
  const itemId = form.querySelector("#hidden-id").value;
  const itemStatus = form.querySelector("#hidden-status").value;
  const itemDescription = form.querySelector("#description").value;
  const formData = new FormData();
  formData.append("id", itemId);
  formData.append("status", itemStatus);
  formData.append("description", itemDescription);
  await fetch("/todo", {
    method: "PUT",
    body: formData,
  });
  fetchItems();
}

async function toggleStatus(item) {
  const itemId = item.id;
  const status = item.dataset.status === "true";
  const description = item.querySelector("span").textContent.trim();
  const formData = new FormData();
  formData.append("id", itemId);
  formData.append("status", !status);
  formData.append("description", description);
  await fetch("/todo", {
    method: "PUT",
    body: formData,
  });
  await fetchItems();
}

function cancelEdit(item) {
  const itemDiv = item.querySelector("div");
  const form = item.querySelector("form");
  itemDiv.classList.remove("hidden");
  form.remove();
}

async function deleteItem(itemDiv) {
  const item = itemDiv.parentElement;
  const itemId = item.id;
  const formData = new FormData();
  formData.append("id", itemId);
  await fetch("/todo", {
    method: "DELETE",
    body: formData,
  });
  fetchItems();
}
