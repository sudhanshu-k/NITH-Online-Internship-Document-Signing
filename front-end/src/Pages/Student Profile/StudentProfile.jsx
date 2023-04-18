import React, { useState } from "react";
import "./StudentProfile.css";
import { useContext } from "react";
import { AppContext } from "../../App";
import { useNavigate } from "react-router-dom";

function StudentProfile() {

	const { userState, setUserState } = useContext(AppContext);
	const navigate = useNavigate();

	const handleLogout = () => {
		setUserState({
			email: "",
			firstname: "",
			lastname: "",
			isfaculty: false,
			isloggedin: false,
			level: "",
		});
		localStorage.removeItem("userState");
		navigate("/");
	};


	const roll = userState.email.slice(0, 8);

	return (
		<div className="student-profile-container">
			<div className="user-container">
				<p className="user-detail">User Name: {userState.firstname + " " + userState.lastname} </p>
				<p className="user-detail">Roll Number: {roll} </p>
				<button onClick={handleLogout}>
					Logout
				</button>
			</div>
		</div>
	);
}

export default StudentProfile;
