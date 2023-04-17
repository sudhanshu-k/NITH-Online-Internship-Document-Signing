import React, { useState } from "react";
import "./StudentProfile.css";
import { useContext } from "react";
import { AppContext } from "../../App";

function StudentProfile() {
	const { userState } = useContext(AppContext);
	const roll = userState.email.slice(0, 8);

	return (
		<div className="student-profile-container">
			<div className="user-container">
				<p className="user-detail">User Name: {userState.firstname + " " + userState.lastname} </p>
				<p className="user-detail">Roll Number: {roll} </p>
			</div>
		</div>
	);
}

export default StudentProfile;
