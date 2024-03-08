import { useEffect, useState } from "react";

function App() {
  const [name, setName] = useState("");
  const [users, setUsers] = useState([]);
  const [editingUser, setEditingUser] = useState(null);

  async function loadUsers() {
    const response = await fetch(import.meta.env.VITE_API + "/users");
    const data = await response.json();
    setUsers(data.users);
  }
  useEffect(() => {
    loadUsers();
  }, []);
  const handleSubmit = async (e) => {
    e.preventDefault();
    await fetch(import.meta.env.VITE_API + "/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name }),
    });
    loadUsers();
  };
  async function deleteUser(id) {
    await fetch(`${import.meta.env.VITE_API}/users/${id}`, {
      method: "DELETE",
    });
    loadUsers();
  }
  const handleEditUser = async (id) => {
    const userToEdit = users.find(user => user._id === id);
    setEditingUser(userToEdit);
};

const handleSaveEdit = async () => {
    await fetch(`${import.meta.env.VITE_API}/users/${editingUser._id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(editingUser),
    }); console.log(editingUser)
    loadUsers();
    setEditingUser(null);
};
  
  return (
    <div className="flex-1 bg-gray-900 text-white p-10 text-xl max-w-3xl">
      <h1 className="text-5xl p-10 text-sky-300">Hello GoLang!</h1>
      <form onSubmit={handleSubmit}>
        <input
          className="rounded-xl p-2 text-black w-1/2 my-2"
          type="name"
          name="name"
          placeholder="Write the user name..."
          onChange={(e) => setName(e.target.value)}
        />
        <button className="ml-4 bg-sky-300 p-2 rounded-xl text-black px-6">
          Save
        </button>
      </form>
      <ul>
        {users != null ? (users.map((user) => (
          <li className="border border-sky-300 rounded-xl p-1 m-2 px-6 flex justify-between" key={user._id}>
            {editingUser && editingUser._id === user._id ? (
    <>
        <div contentEditable={true} onBlur={(e) => setEditingUser({ ...editingUser, Name: e.target.textContent })}>
            {user.Name}
        </div>
        <button onClick={handleSaveEdit}>Save</button>
    </>
) : (
    <>
        {user.Name}
        <div className="flex gap-4"> 
            <button onClick={() => handleEditUser(user._id)}>Edit</button>
            <button onClick={() => deleteUser(user._id)}>Delete</button>
        </div>
    </>
)}

          </li>
        ))):( <li>No users in the DataBase</li>) }
      </ul>
    </div>
  );
}

export default App;
