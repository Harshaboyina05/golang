import React, { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import axios from 'axios';
import "./Payment.css";
import { useAuth } from 'react-oidc-context';

const Payment = () => {
    const location = useLocation();
    const { selectedCar, userDetails } = location.state || {};

    const getCurrentTime = () => {
        return new Date().getHours() < 12 ? '15' : '200';
    };

    useEffect(() => {
        console.log(selectedCar, userDetails, "paymenttttttttt")
    })
    const auth=useAuth();
    const [paymentMethod, setPaymentMethod] = useState('Card'); // Default payment method
    const [amount, setAmount] = useState(getCurrentTime());

    const handlePaymentConfirmation = async () => {
        try {
            // Perform payment processing logic here based on the selected payment method
            // For demonstration purposes, I'm using a simple alert
            alert(`Payment Successful! Amount: $${amount}, Payment Method: ${paymentMethod}`);

            // Update the status to "completed" in the backend
            const paymentData = {
                amount,
                payment_method: paymentMethod,
                status: 'completed'
            };

            // Make a POST request to the payment API endpoint to add payment details
            const paymentResponse = await axios.post('http://localhost:4002/payment',{ headers: {
                Authorization: `Bearer ${auth.user.access_token}`,
                'Content-Type': 'application/json',
              }}, paymentData);

            console.log('Payment record added successfully:', paymentResponse.data);
        } catch (error) {
            console.error('Error confirming payment:', error);
        }
    };

    const handleChangePaymentMethod = (e) => {
        setPaymentMethod(e.target.value);
    };

    return (
        <div className="center">
         <h2>Confirm your Car Booking</h2>   
            <div className="details-container">
                <div className="user-details">
                    <h3>User Details:</h3>
                    <p>First Name: {userDetails.first_name}</p>
                    <p>Last Name: {userDetails.last_name}</p>
                    <p>Email: {userDetails.email}</p>
                    <p>Phone Number: {userDetails.phone_number}</p>
                    <p>Address: {userDetails.address}</p>
                </div>
                <div className="car-details">
                    <h3>Selected Car Details:</h3>
                    <p>Make: {selectedCar.selectedCar.make}</p>
                    <p>Model: {selectedCar.selectedCar.model}</p>
                    <p>License Plate: {selectedCar.selectedCar.license_plate}</p>
                    <p>Color: {selectedCar.selectedCar.color}</p>
                    <p>Mileage: {selectedCar.selectedCar.mileage}</p>
                    <p>Status: {selectedCar.selectedCar.status}</p>
                </div>
            </div>
            <h3>Payment Options:</h3>
            <p>Amount: ${amount}</p>
            <div>
                <label htmlFor="paymentMethod">Payment Method:</label>
                <select id="paymentMethod" value={paymentMethod} onChange={handleChangePaymentMethod}>
                    <option value="Card">Card</option>
                    <option value="UPI">UPI</option>
                </select>
            </div>
            <div className="button-container">
                <button type="submit" className="button" onClick={handlePaymentConfirmation}>Confirm Payment</button>
            </div>
        </div>
    );
};

export default Payment;
