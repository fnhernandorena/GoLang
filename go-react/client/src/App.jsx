import { useEffect, useState } from "react";

function App() {
  const [name, setName] = useState("");
  const [users, setUsers] = useState([]);

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
  return (
    <div className="flex-1 bg-gray-900 text-white p-10">
      <h1 className="text-5xl p-10 text-sky-300">Hello GoLang!</h1>
      <form onSubmit={handleSubmit}>
        <input className="rounded-xl p-2 text-black w-1/2 my-2"
          type="name"
          name="name"
          placeholder="Write the user name..."
          onChange={(e) => setName(e.target.value)}
        />
        <button className="ml-4 bg-sky-300 p-2 rounded-xl text-black px-6">Save</button>
      </form>
      <ul>
        {users!=null&&users.map((user) => (
          <li className="border border-sky-300 rounded-xl p-1 m-1 px-3" key={user._id}>
            {user.Name}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
