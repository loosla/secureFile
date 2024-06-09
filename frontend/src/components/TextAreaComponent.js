import React, { useEffect, useState } from 'react';

const TextAreaComponent = () => {
  const [textAreaValue, setTextAreaValue] = useState('');

  useEffect(() => {
    fetch('http://localhost:8080/api/text')
      .then(response => response.json())
      .then(data => setTextAreaValue(data.text))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  const handleSave = () => {
    fetch('http://localhost:8080/api/save', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ text: textAreaValue }),
    })
      .then(response => {
        if (response.ok) {
          alert('Text saved successfully');
        } else {
          alert('Failed to save text');
        }
      })
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
