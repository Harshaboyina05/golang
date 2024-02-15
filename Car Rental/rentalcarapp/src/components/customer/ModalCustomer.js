import React, { useState } from 'react';

const Modal = ({ isOpen, onClose, onSubmit }) => {
  const initialData = {
    first_name: '',
    last_name: '',
    email: '',
    phone_number: '',
    address: '', // Set a default value for status
    created_at: new Date().toISOString(), // Set current timestamp for created_at
    updated_at: new Date().toISOString(), // Set current timestamp for updated_at
  };

  const [formData, setFormData] = useState(initialData);

  const handleInputChange = (fieldName, value) => {
    setFormData(prevData => ({
      ...prevData,
      [fieldName]: value,
    }));
  };

  const handleSubmit = async () => {
    const dataPayload = {
      first_name: formData.first_name,
      last_name: formData.last_name,
      email: formData.email, // Convert to integer
      phone_number: formData.phone_number,
      address: formData.address,
      created_at: formData.created_at,
      updated_at: formData.updated_at
    };
    onSubmit(dataPayload);
    setFormData(initialData);
    onClose();
  };

  if (!isOpen) return null;


  return (
    <div
      style={{
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        background: 'rgba(0, 0, 0, 0.5)',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
      }}
    >
      <div
        style={{
          position: 'relative',
          background: 'white',
          width: 500,
          height: 'auto',
          margin: 'auto',
          padding: '4%',
          border: '2px solid #000',
          borderRadius: '10px',
          boxShadow: '2px solid black',
        }}
      >
        <button
          onClick={onClose}
          style={{
            position: 'absolute',
            top: 5,
            right: 4,
            background: 'none',
            border: 'none',
            cursor: 'pointer',
            padding: 0,
            fontSize: '20px',
          }}
        >
          <span
            style={{
              display: 'inline-block',
              width: '30px',
              height: '30px',
              borderRadius: '50%',
              background: '#eee',
              textAlign: 'center',
              lineHeight: '30px',
            }}
          >
            &times;
          </span>
        </button>

        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            marginBottom: '10px',
          }}
        >
          {Object.entries(formData).map(([fieldName, value]) => (
            <div
              key={fieldName}
              style={{ marginBottom: '10px', display: 'flex' }}
            >
              <label
                style={{
                  color: 'black',
                  paddingRight: '20px',
                  width: '120px',
                  flex: '0 0 120px',
                  boxSizing: 'border-box',
                  marginTop: '1px',
                }}
              >
                {fieldName}:
              </label>
              <input
                type="text"
                value={value}
                onChange={e =>
                  handleInputChange(fieldName, e.target.value)
                }
                style={{
                  paddingLeft: '5px',
                  flex: 1,
                  height: '30px',
                  borderRadius: '5px',
                  border: '1px solid #ccc',
                }}
              />
            </div>
          ))}
        </div>

        <div style={{ textAlign: 'center' }}>
          <button
            onClick={handleSubmit}
            style={{
              backgroundColor: '#4CAF50',
              color: 'white',
              padding: '10px',
              border: 'none',
              borderRadius: '5px',
              cursor: 'pointer',
            }}
          >
            Submit
          </button>
        </div>
      </div>
    </div>
  );
};

export default Modal;
