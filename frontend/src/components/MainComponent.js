import React, { useState, useRef } from 'react';

const TextAreaComponent = () => {
  const [textAreaValue, setTextAreaValue] = useState('');
  const [password, setPassword] = useState('');
  const [file, handleFileUpload] = useState('');

  const fileInputRef = useRef(null);

  const openFileDialog = () => {
    fileInputRef.current.click();
  };

  const handleSave = () => {
    window.api.filesSave({ password: password, content: textAreaValue })
      .then(() => alert('Text saved successfully'))
      .catch(error => console.error('Error saving data:', error));
  };

  const handleFetchContent = async () => {
    const content = await window.api.filesContent({file: file, password:password});
    setTextAreaValue(content);
  };

  return (
    <div>
      <div class="container">
        <button onClick={openFileDialog}>Choose File</button>
        <input
          type="file"
          ref={fileInputRef}
          style={{ display: 'none' }}
          onChange={(e) => handleFileUpload(e.target.value)}
        />
        <div>
          <label htmlFor="password">Password: </label>
          <input
            id="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Enter password"
          />
        </div>
      </div>
      <br />
      <button onClick={handleFetchContent}>Decrypt</button>
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
