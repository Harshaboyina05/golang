import React, { useState, useEffect } from 'react';
import Modal from './Modal';
import { MdDeleteOutline } from 'react-icons/md';
import { useAuth } from 'react-oidc-context';

const NotesList = ({ notesApp }) => {
  const [notes, setNotes] = useState([]);
  const [showAddNotePopup, setShowAddNotePopup] = useState(false);
  const [openModal, setOpenModal] = useState(false);
  const [newNote, setNewNote] = useState({
    make: '',
    model: '',
    year: '',
    license_plate: '',
    color: '',
    mileage: '',
    status: 'available',
    created_at: new Date().toISOString(),
    updated_at: new Date().toISOString(),
  });
  const auth = useAuth();

  useEffect(() => {
    // Fetch notes from the backend API
    const fetchNotes = async () => {
      try {
        let envString = 'REACT_APP_MICROSERVICE_CARMANAGEMENT'
        const response = await fetch(process.env[envString] + `/car`, {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${auth.user.access_token}`,
            'Content-Type': 'application/json',
          },
        });
        const data = await response.json();
        if (data != null) setNotes(data);
      } catch (error) {
        console.error('Error fetching notes:', error);
      }
    };
    fetchNotes();
  }, []);

  const handleInputChange = e => {
    const { name, value } = e.target;
    setNewNote(prevNote => ({
      ...prevNote,
      [name]: value,
    }));
  };

  const handlePopupClose = () => {
    setShowAddNotePopup(false);
    setNewNote({
      make: '',
      model: '',
      year: '',
      license_plate: '',
      color: '',
      mileage: '',
      status: 'available',
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    });
  };

  const handleAddNote = async (data) => {
    try {
      const response = await fetch('http://localhost:4000/car', {
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
        setNotes(prevNotes => [...prevNotes, data]);
        setNewNote({
          make: '',
          model: '',
          year: '',
          license_plate: '',
          color: '',
          mileage: '',
          status: 'available',
          created_at: new Date().toISOString(),
          updated_at: new Date().toISOString(),
        });
      } else {
        console.error('Failed to add note:', response.statusText);
      }
    } catch (error) {
      console.error('Error adding note:', error);
    }
  };
  
  const handleDeleteNote = async id => {
    try {
      const response = await fetch(`http://localhost:4000/car/${id}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,

          'Content-Type': 'application/json',
        },
      });
  
      if (response.ok) {
        setNotes(prevNotes => prevNotes.filter(note => note.id !== id));
      } else {
        console.error('Failed to delete note:', response.statusText);
      }
    } catch (error) {
      console.error('Error deleting note:', error);
    }
  };
  

  return (
    <div className="container ping">
      <button className="ping-button" onClick={() => setOpenModal(true)} style={{ alignItems: 'revert-layer' }}>
        Add Car
      </button>
      <Modal isOpen={openModal} onClose={() => setOpenModal(false)} onSubmit={handleAddNote} />
      {/* Your popup JSX code */}
      <table style={{ borderCollapse: 'collapse', width: '100%' }}>
        <thead>
          <tr style={{ color: 'black', backgroundColor: '#f2f2f2' }}>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>id</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>make</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>model</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>year</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>license_plate</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>color</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>mileage</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>status</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>created_at</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>updated_at</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Actions</th>
          </tr>
        </thead>
        <tbody>
          {notes.map((note, index) => (
            <tr key={index} style={{ color: 'black', backgroundColor: index % 2 === 0 ? '#f9f9f9' : 'white' }}>
              {/* Render your note data here */}
              <td>{index + 1}</td>
              {/* Render other columns */}
              <td>{note.make}</td>
              <td>{note.model}</td>
              <td>{note.year}</td>
              <td>{note.license_plate}</td>
              <td>{note.color}</td>
              <td>{note.mileage}</td>
              <td>{note.status}</td>
              <td>{note.created_at}</td>
              <td>{note.updated_at}</td>
              <td>
                <MdDeleteOutline
                  style={{ fontSize: 20, color: 'red', cursor: 'pointer' }}
                  onClick={() => handleDeleteNote(note.id)}
                />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default NotesList;
