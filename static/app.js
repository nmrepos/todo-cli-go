let current = [];

async function load() {
    const resp = await fetch('/todos');
    const tasks = await resp.json();
    current = tasks;
    const ul = document.getElementById('list');
    ul.innerHTML = '';
    tasks.forEach(t => {
        const li = document.createElement('li');
        li.textContent = (t.Done ? '[x] ' : '[ ] ') + t.Title;
        const doneBtn = document.createElement('button');
        doneBtn.textContent = 'Done';
        doneBtn.onclick = () => completeTask(t.ID, t.Title);
        const delBtn = document.createElement('button');
        delBtn.textContent = 'Delete';
        delBtn.onclick = () => deleteTask(t.ID);
        li.appendChild(doneBtn);
        li.appendChild(delBtn);
        ul.appendChild(li);
    });
}

async function addTask() {
    const text = document.getElementById('task').value;
    await fetch('/todos', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title: text })
    });
    document.getElementById('task').value = '';
    load();
}

async function completeTask(id, title) {
    await fetch('/todos/' + id, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title: title, done: true })
    });
    load();
}

async function deleteTask(id) {
    await fetch('/todos/' + id, { method: 'DELETE' });
    load();
}

window.onload = load;
