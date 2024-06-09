import React, { useEffect, useState } from 'react';
import './App.css';
import TextAreaComponent from './components/TextAreaComponent';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetch('http://localhost:8080/api/hello')
      .then(response => response.text())
      .then(data => setMessage(data));
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <TextAreaComponent />
        <p>{message}</p>
      </header>
    </div>
  );
}

export default App;
