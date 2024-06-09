import React, { useEffect, useState } from 'react';

const TextAreaComponent = () => {
  const [textAreaValue, setTextAreaValue] = useState('');

  useEffect(() => {
    window.api.fetchText()
      .then(data => setTextAreaValue(data.text))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  const handleSave = () => {
    window.api.saveText({ text: textAreaValue })
      .then(() => alert('Text saved successfully'))
      .catch(error => console.error('Error saving data:', error));
  };

  return (
    <div>
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
