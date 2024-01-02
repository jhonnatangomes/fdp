import { useEffect, useState } from 'react';

export default function App() {
  const [name, setName] = useState('');
  const [serverMessage, setServerMessage] = useState('');
  const [socket, setSocket] = useState<WebSocket | null>(null);
  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8080');
    setSocket(socket);
    socket.addEventListener('close', event => {
      console.log('close', event);
    });
    socket.addEventListener('message', event => {
      setServerMessage(event.data);
    });
  }, []);
  return (
    <div className='h-screen w-full flex flex-col gap-y-4 justify-center items-center'>
      <span>Coloque seu nome</span>
      <input
        className='border border-solid border-gray-300'
        value={name}
        onChange={onChange}
      />
      <button className='bg-gray-200 p-2 rounded' onClick={joinGame}>
        Entrar no jogo
      </button>
      {serverMessage && <span>{serverMessage}</span>}
    </div>
  );
  function onChange(e: React.ChangeEvent<HTMLInputElement>) {
    setName(e.target.value);
  }
  function joinGame() {
    socket?.send('join');
  }
}
