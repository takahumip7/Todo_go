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

  return (
    <div style={{ maxWidth: '600px', margin: '0 auto', padding: '16px' }}>
      <h1 style={{ fontSize: '28px', fontWeight: 'bold', marginBottom: '24px' }}>
        Todo List
      </h1>
      <ul style={{ listStyle: 'none', padding: 0 }}>
        {todos.map((todo) => (
          <li
            key={todo.id}
            style={{
              display: 'flex',
              justifyContent: 'space-between',
              alignItems: 'center',
              padding: '12px 16px',
              marginBottom: '8px',
              borderRadius: '8px',
              backgroundColor: statusColors[todo.status],
              boxShadow: '0 1px 3px rgba(0,0,0,0.1)',
            }}
          >
            <span>
              {todo.title} <br />
              <small>{new Date(todo.created_at).toLocaleString()}</small>
            </span>
            <span style={{ fontWeight: 'bold' }}>
              {statusLabels[todo.status]}
            </span>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
