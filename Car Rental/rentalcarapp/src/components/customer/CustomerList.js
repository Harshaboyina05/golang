import React, { useState, useEffect } from 'react';
import Modal from './ModalCustomer';
import { MdDeleteOutline } from 'react-icons/md';
import { useAuth } from 'react-oidc-context';

const CustomerList = ({ customerApp }) => {
  const [customer, setCustomer] = useState([]);
  const [showAddCustomerPopup, setShowAddCustomerPopup] = useState(false);
  const [openModal, setOpenModal] = useState(false);
  const [newCustomer, setNewCustomer] = useState({ 
    first_name: '',
    last_name: '',
    email: '',
    phone_number: '',
    address: '',
    created_at: new Date().toISOString(),
    updated_at: new Date().toISOString(),
  });
  const auth = useAuth();

  useEffect(() => {
    // Fetch Customers from the backend API
    const fetchCustomers = async () => {
      try {
        let envString = 'REACT_APP_MICROSERVICE_CUSTOMERMANAGEMENT'
        const response = await fetch(process.env[envString] + `/customer`, {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${auth.user.access_token}`,
            'Content-Type': 'application/json',
          },
        });
        const data = await response.json();
        if (data != null) setCustomer(data);
      } catch (error) {
        console.error('Error fetching Customers:', error);
      }
    };
    fetchCustomers();
  }, []);

  const handleInputChange = e => {
    const { name, value } = e.target;
    setNewCustomer(prevCustomer => ({
      ...prevCustomer,
      [name]: value,
    }));
  };

  const handlePopupClose = () => {
    setShowAddCustomerPopup(false);
    setNewCustomer({
      first_name: '',
      last_name: '',
      email: '',
      phone_number: '',
      address: '',
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    });
  };

  const handleAddCustomer = async (data) => {
    try {
      const response = await fetch('http://localhost:4001/customer', {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,

          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
  
      if (response.ok) {
        const data = await response.json();
        console.log(data)
        setCustomer(prevCustomers => [...prevCustomers, data]);
        setNewCustomer({
          first_name: '',
          last_name: '',
          email: '',
          phone_number: '',
          address: '',
          created_at: new Date().toISOString(),
          updated_at: new Date().toISOString(),
        });
      } else {
        console.error('Failed to add Customer:', response.statusText);
      }
    } catch (error) {
      console.error('Error adding Customer:', error);
    }
  };
  
  const handleDeleteCustomer = async id => {
    try {
      const response = await fetch(`http://localhost:4001/customer/${id}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,

          'Content-Type': 'application/json',
        },
      });
  
      if (response.ok) {
        setCustomer(prevCustomers => prevCustomers.filter(Customer => Customer.id !== id));
      } else {
        console.error('Failed to delete Customer:', response.statusText);
      }
    } catch (error) {
      console.error('Error deleting Customer:', error);
    }
  };
  

  return (
    <div className="container ping">
      <button className="ping-button" onClick={() => setOpenModal(true)} style={{ alignItems: 'revert-layer' }}>
        Add customer
      </button>
      <Modal isOpen={openModal} onClose={() => setOpenModal(false)} onSubmit={handleAddCustomer} />
      {/* Your popup JSX code */}
      <table style={{ borderCollapse: 'collapse', width: '100%' }}>
        <thead>
          <tr style={{ color: 'black', backgroundaddress: '#f2f2f2' }}>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>id</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>first_name</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>last_name</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>email</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>phone_number</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>address</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Actions</th>
          </tr>
        </thead>
        <tbody>
          {customer.map((Customer, index) => (
            <tr key={index} style={{ color: 'black', backgroundaddress: index % 2 === 0 ? '#f9f9f9' : 'white' }}>
              {/* Render your Customer data here */}
              <td>{index + 1}</td>
              {/* Render other columns */}
              <td>{Customer.first_name}</td>
              <td>{Customer.last_name}</td>
              <td>{Customer.email}</td>
              <td>{Customer.phone_number}</td>
              <td>{Customer.address}</td>
              <td>
                <MdDeleteOutline
                  style={{ fontSize: 20, address: 'red', cursor: 'pointer' }}
                  onClick={() => handleDeleteCustomer(Customer.id)}
                />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default CustomerList;
