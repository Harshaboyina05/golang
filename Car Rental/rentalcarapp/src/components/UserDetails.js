import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './UserDetails.css'; // Import the CSS file
import axios from 'axios';
import { useAuth } from 'react-oidc-context';
import { useLocation } from 'react-router-dom';

const UserDetails = () => { // Receive selectedCar as props
const location = useLocation();
  const selectedCar = location.state
  const [formData, setFormData] = useState({
    first_name: '',
    last_name: '',
    email: '',
    phone_number: '',
    address: '',
  });
  const navigate = useNavigate();
  const auth = useAuth();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };
  useEffect(()=>{
    console.log(selectedCar,'carrrrr')
  })

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:4001/customer', formData, {
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
      });
      console.log('User details added to the database:', response.data);
      console.log(selectedCar,formData)
      navigate('/payment',{ state: { selectedCar, userDetails: formData } }); // Pass userDetails as state
    } catch (error) {
      console.error('Error adding user details:', error);
    }
  };

  return (
    <div className="center user-details">
      <h2>Add your details here</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label className="label">First Name:</label>
          <input type="text" name="first_name" value={formData.first_name} onChange={handleChange} className="input" />
        </div>
        <div>
          <label className="label">Last Name:</label>
          <input type="text" name="last_name" value={formData.last_name} onChange={handleChange} className="input" />
        </div>
        <div>
          <label className="label">Email:</label>
          <input type="email" name="email" value={formData.email} onChange={handleChange} className="input" />
        </div>
        <div>
          <label className="label">Phone Number:</label>
          <input type="text" name="phone_number" value={formData.phone_number} onChange={handleChange} className="input" />
        </div>
        <div>
          <label className="label">Address:</label>
          <input type="text" name="address" value={formData.address} onChange={handleChange} className="input" />
        </div>
        <div className="button-container">
          <button type="submit" className="button">Proceed with Payment</button>
        </div>
      </form>
    </div>
  );
};

export default UserDetails;
