import {useEffect, useState} from 'react';

// ステータスをフロント側（enum）で管理
enum TodoStatus {
  未着手 = 0,
  完了 = 1
}

interface Todo {
  id: number;
  title: string;
  status: TodoStatus;
  created_at: string;
}

// completed(boolean) → status(enum) に変換
const mapCompletedToStatus = (completed: boolean): TodoStatus => {
  return completed ? TodoStatus.完了 : TodoStatus.未着手;
};

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTitle, setNewTitle] = useState(""); // 入力フォーム用

  const statusLabels = ['未着手', '完了'];
  const statusColors = ['#ffe5e5', '#d1ffd6'];

  // GET で Todo 一覧を取得
  useEffect(() => {
    fetch('http://localhost:8080/todos')
      .then(res => res.json())
      .then((data: {id: number; title: string; completed: boolean; created_at: string}[]) =>{
        // completed → status に変換して state にセット
        const mapped = data.map(todo => ({
          ...todo,
          status: mapCompletedToStatus(todo.completed),
        }));
        setTodos(mapped);
      })
      .catch(err => console.error(err));
  }, []);

  // POST (新規作成)
  const addTodo = async () => {
    if (!newTitle.trim()) return; // 空文字は無視

    const res = await fetch("http://localhost:8080/todos", {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({ title: newTitle, completed: false}),
    });

    const newTodo = await res.json();

    // 追加した Todo を state に反映
    setTodos([
      ...todos,
      { ...newTodo, status: mapCompletedToStatus(newTodo.completed)},
    ]);

    setNewTitle(""); // 入力欄クリア
  }

  // DELETE(削除)
  const deleteTodo = async (id: number) => {
    if (!window.confirm("このタスクを削除しますか？")) return;

    const res = await fetch(`http://localhost:8080/todos/${id}`, {
      method: "DELETE",
    });

    if (res.ok) {
      //stateから削除
      setTodos(todos.filter(todo => todo.id !== id));
    } else {
      const errText = await res.text();
      alert("削除に失敗しました。: " + errText);
    }
  };


  return (
    <div style={{ maxWidth: "600px", margin: "0 auto", padding: "16px" }}>
      <h1 style={{ fontSize: "28px", fontWeight: "bold", marginBottom: "24px" }}>
        Todo List
      </h1>

      {/* 新規作成フォーム */}
      <div
        style={{
          display: "flex",
          gap: "8px",
          marginBottom: "20px",
        }}
      >
        <input
          type="text"
          value={newTitle}
          onChange={(e) => setNewTitle(e.target.value)}
          placeholder="新しいタスクを入力"
          style={{
            flex: 1,
            padding: "8px",
            border: "1px solid #ccc",
            borderRadius: "6px",
          }}
        />
        <button
          onClick={addTodo}
          style={{
            padding: "8px 12px",
            backgroundColor: "#4CAF50",
            color: "white",
            border: "none",
            borderRadius: "6px",
            cursor: "pointer",
          }}
        >
          追加
        </button>
      </div>

      <ul style={{ listStyle: "none", padding: 0 }}>
        {todos.map((todo) => (
          <li
            key={todo.id}
            style={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
              padding: "12px 16px",
              marginBottom: "8px",
              borderRadius: "8px",
              backgroundColor: statusColors[todo.status],
              boxShadow: "0 1px 3px rgba(0,0,0,0.1)",
            }}
          >
            <div>
              {todo.title} <br />
              <small>{new Date(todo.created_at).toLocaleString()}</small>
            </div>
            <div style={{ display: "flex", alignItems: "center", gap: "8px"}}>
              <span style={{ fontWeight: "bold" }}>
                {statusLabels[todo.status]}
              </span>
              {/* 🗑 削除ボタン */}
              <button onClick={() => deleteTodo(todo.id)} 
                style={{ padding: "4px 8px", backgroundColor: "#f44336", color: "white", border: "none", borderRadius: "6px", cursor: "pointer"}}>
                削除
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
