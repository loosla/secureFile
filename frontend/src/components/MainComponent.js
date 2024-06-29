import React, { useState } from 'react';

const TextAreaComponent = () => {
  const [textAreaValue, setTextAreaValue] = useState('');
  const [password, setPassword] = useState('');

  const handleUpdate = () => {
    window.api.update({ password: password, content: textAreaValue })
      .then(() => alert('Text saved successfully'))
      .catch(error => console.error('Error saving data:', error));
  };

  const handleGet = async () => {
    const content = await window.api.get(password);
    setTextAreaValue(content);
  };

  return (
    <div>
      <label htmlFor="password">Password: </label>
      <input
        id="password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Enter password"
      />
      <br />
      <button onClick={handleGet}>Decrypt</button>
      <br />
      <textarea
        value={textAreaValue}
        onChange={(e) => setTextAreaValue(e.target.value)}
        rows="10"
        cols="50"
      />
      <br />
      <button onClick={handleUpdate}>Save</button>
    </div>
  );
};

export default TextAreaComponent;
