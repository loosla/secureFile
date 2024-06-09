import React, { useState } from 'react';

const TextAreaComponent = () => {
  const [textAreaValue, setTextAreaValue] = useState('');

  const handleChange = (event) => {
    setTextAreaValue(event.target.value);
  };

  return (
    <div>
      <textarea
        value={textAreaValue}
        onChange={handleChange}
        placeholder="Enter text here"
        rows="10"
        cols="50"
      />
      <p>You entered: {textAreaValue}</p>
    </div>
  );
};

export default TextAreaComponent;
