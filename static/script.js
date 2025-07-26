function createTodo() {
  const input_data = document.querySelector("#new-item-form > input");

  fetch("https://dummyjson.com/todos/add", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      todo: input_data.value,
      completed: false,
      userId: 1,
    }),
  })
    .then((res) => res.json())
    .then(console.log)
    .catch(console.error);
}

function formTemplate(currentText) {
  const form = document.createElement("form");
  form.classList.add("edit-form");
  form.innerHTML = `
      <input type="text" class="edit-form-input" value="${currentText}">
      <button type="button" class="task-list-button" onclick="">
        Save
      </button>
      <button type="button" class="task-list-button" onclick="cancelEdit(this.parentElement.parentElement)">
        Cancel
      </button>
  `;
  return form;
}

function editItem(itemDiv) {
  // const itemSpan = document.querySelector(`#${itemId} > span`);
  const currentText = itemDiv.querySelector("span").textContent.trim();
  const form = formTemplate(currentText);
  itemDiv.parentElement.appendChild(form);
  itemDiv.classList.add("hidden");
}

function cancelEdit(item) {
  const itemDiv = item.querySelector("div");
  const form = item.querySelector("form");
  itemDiv.classList.remove("hidden");
  form.remove();
}
