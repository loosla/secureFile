import React, { useState } from 'react';

const TextAreaComponent = () => {
  const [textAreaValue, setTextAreaValue] = useState('');
  const [password, setPassword] = useState('');

  const handleSave = () => {
    window.api.saveText({ text: textAreaValue })
      .then(() => alert('Text saved successfully'))
      .catch(error => console.error('Error saving data:', error));
  };

  const handleFetchFile = async () => {
    const content = await window.api.fetchFile(password);
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
      <button onClick={handleFetchFile}>Decrypt</button>
      <br />
      <textarea
        value={textAreaValue}
        onChange={(e) => setTextAreaValue(e.target.value)}
        rows="10"
        cols="50"
      />
      <br />
      <button onClick={handleSave}>Save</button>
    </div>
  );
};

export default TextAreaComponent;
