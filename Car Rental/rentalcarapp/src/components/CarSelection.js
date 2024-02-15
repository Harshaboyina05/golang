// CarSelection.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useAuth } from 'react-oidc-context';
import { useNavigate } from 'react-router-dom'; // Updated import
import './CarSelection.css';

const CarSelection = () => {
  const [cars, setCars] = useState([]);
  const [selectedCar, setSelectedCar] = useState(null);
  const auth = useAuth();
  const navigate = useNavigate(); // Updated hook

  useEffect(() => {
    // Fetch cars from the backend API when the component mounts
    const fetchData = async () => {
      try {
        const response = await axios.get('http://localhost:4000/car', {
          headers: {
            Authorization: `Bearer ${auth.user.access_token}`,
            'Content-Type': 'application/json',
          }
        });
        setCars(response.data);
      } catch (error) {
        console.error('Error fetching cars:', error);
      }
    };
    fetchData();
  }, []);

  const handleCarSelect = (car) => {
    setSelectedCar(car);
  };

  const handleBookNow = () => {
    navigate('/booking', { state: { selectedCar } }); // Pass selectedCar as state
  };

  return (
    <div className="center">
      <h2>Car Selection</h2>
      <p>Please select an available car:</p>
      <select onChange={(e) => handleCarSelect(JSON.parse(e.target.value))}>
        <option value="">Select a car</option>
        {cars.filter(car => car.status === 'available').map((car) => (
          <option key={car.id} value={JSON.stringify(car)}>
            {car.make} {car.model} - {car.license_plate} - {car.color} - {car.mileage} - {car.status}
          </option>
        ))}
      </select>
      {selectedCar && (
        <div>
          <h3>Selected Car Details</h3>
          <p>Make: {selectedCar.make}</p>
          <p>Model: {selectedCar.model}</p>
          <p>License Plate: {selectedCar.license_plate}</p>
          <p>Color: {selectedCar.color}</p>
          <p>Mileage: {selectedCar.mileage}</p>
          <p>Status: {selectedCar.status}</p>
          <div className="button-container">
              <button className="book-now-button" onClick={handleBookNow}>Book Now</button>
          </div>
        </div>
      )}
    </div>
  );
};

export default CarSelection;
