import React, { useEffect, useState } from 'react';

const TextAreaComponent = () => {
  const [textAreaValue, setTextAreaValue] = useState('');
  const [inputValue, setInputValue] = useState('');


  useEffect(() => {
    window.api.fetchText()
      .then(data => setTextAreaValue(data.text))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  const handleChange = (e) => {
    setInputValue(e.target.value);
  };

  const handleSave = () => {
    window.api.saveText({ text: textAreaValue })
      .then(() => alert('Text saved successfully'))
      .catch(error => console.error('Error saving data:', error));
  };

  const handleDecrypt = () => {
    window.api.saveText({ text: "textAreaValue" })
      .then(() => alert('Text saved successfully'))
      .catch(error => console.error('Error saving data:', error));
  };


  return (
    <div>
      <label htmlFor="hidden-input">Password: </label>
      <input
        id="hidden-input"
        type="password"
        value={inputValue}
        onChange={handleChange}
      />
      <br />
      <button onClick={handleDecrypt}>Decrypt</button>
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
