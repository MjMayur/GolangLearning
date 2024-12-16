import React, { useState } from "react";
import {
  Card,
  CardBody,
  Form,
  FormGroup,
  Label,
  Input,
  Button,
} from "reactstrap";

function UserForm() {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    phone: "",
    message: "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("Form Data Submitted:", formData);
    // Here you can add your form submission logic
  };
  const [records, setRecords] = useState([
    {
      id: 1,
      name: "John Doe",
      email: "john@example.com",
      phone: "123-456-7890",
    },
    {
      id: 2,
      name: "Jane Smith",
      email: "jane@example.com",
      phone: "987-654-3210",
    },
    // Add more records as needed
  ]);

  const handleEdit = (id) => {
    // Implement edit logic here
    console.log("Edit record with id:", id);
  };

  const handleDelete = (id) => {
    const updatedRecords = records.filter((record) => record.id !== id);
    setRecords(updatedRecords);
  };

  const submitForm = async () => {
    // API endpoint
    const apiUrl = "http://localhost:8000/user/create/";

    // Fetch API
    let headers = new Headers();
    headers.append("Content-Type", "application/json");
    fetch(apiUrl, {
      method: "POST",
      headers: headers,
      body: JSON.stringify(formData),
    })
      .then((res) => res.json())
      .then((res) => {
        console.log(res);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <>
      <Card
        style={{
          maxWidth: "500px",
          margin: "auto",
          marginTop: "3rem",
          padding: "20px",
          borderRadius: "10px",
          boxShadow: "0 2px 10px rgba(0,0,0,0.1)",
        }}
      >
        <CardBody className="p-2">
          <h2 style={{ textAlign: "center", marginBottom: "2rem" }}>
            Contact Us
          </h2>
          <Form onSubmit={handleSubmit}>
            <FormGroup>
              <Input
                type="text"
                name="name"
                id="name"
                placeholder="Enter your name"
                value={formData.name}
                onChange={handleChange}
                required
              />
            </FormGroup>
            <FormGroup>
              <Input
                type="email"
                name="email"
                id="email"
                placeholder="Enter your email"
                value={formData.email}
                onChange={handleChange}
                required
              />
            </FormGroup>
            <FormGroup>
              <Input
                type="tel"
                name="phone"
                id="phone"
                placeholder="Enter your phone number"
                value={formData.phone}
                onChange={handleChange}
                required
              />
            </FormGroup>
            <FormGroup>
              <Input
                type="textarea"
                name="message"
                id="message"
                placeholder="Enter your message"
                value={formData.message}
                onChange={handleChange}
                required
              />
            </FormGroup>
            <Button
              color="primary"
              style={{ width: "100%" }}
              onClick={submitForm}
            >
              Submit
            </Button>
          </Form>
        </CardBody>
      </Card>
      <div
        style={{ padding: "20px", display: "flex", justifyContent: "center" }}
      >
        <table
          style={{
            width: "80%",
            borderCollapse: "collapse",
            boxShadow: "0 2px 10px rgba(0,0,0,0.1)",
          }}
        >
          <thead>
            <tr style={{ backgroundColor: "#f2f2f2" }}>
              <th
                style={{
                  padding: "10px",
                  border: "1px solid #dddddd",
                  textAlign: "center",
                }}
              >
                Name
              </th>
              <th
                style={{
                  padding: "10px",
                  border: "1px solid #dddddd",
                  textAlign: "center",
                }}
              >
                Email
              </th>
              <th
                style={{
                  padding: "10px",
                  border: "1px solid #dddddd",
                  textAlign: "center",
                }}
              >
                Phone
              </th>
              <th
                style={{
                  padding: "10px",
                  border: "1px solid #dddddd",
                  textAlign: "center",
                }}
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            {records.map((record) => (
              <tr key={record.id} style={{ backgroundColor: "#ffffff" }}>
                <td style={{ padding: "10px", border: "1px solid #dddddd" }}>
                  {record.name}
                </td>
                <td style={{ padding: "10px", border: "1px solid #dddddd" }}>
                  {record.email}
                </td>
                <td style={{ padding: "10px", border: "1px solid #dddddd" }}>
                  {record.phone}
                </td>
                <td style={{ padding: "10px", border: "1px solid #dddddd" }}>
                  <button
                    onClick={() => handleEdit(record.id)}
                    style={{ marginRight: "10px" }}
                  >
                    Edit
                  </button>
                  <button onClick={() => handleDelete(record.id)}>
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}

export default UserForm;
