import React, { useState } from "react";
import "./StudentProfile.css";

function StudentProfile() {
	const user = {
		userName: "Student Name",
		img: "http://shorturl.at/DVYZ6",
		rollNumber: 20,
		email: "20xxx000@nith.ac.in",
		phone: 9999999999,
	};
	const [userState, setUserState] = useState(user);
	const [thisName, setThisName] = useState("");
	const [isVisibleName, setIsVisibleName] = useState(false);
	const [isVisibleRoll, setIsVisibleRoll] = useState(false);

	const handleClick = () => {
		if (thisName.length === 0) {
			alert();
			return;
		}
		setUserState((prev) => {
			return {
				...prev,
				userName: thisName,
			};
		});
		console.log(user.userName);
		setIsVisibleName(false);
	};

	return (
		<div className="student-profile-container">
			<div className="user-container">
				<img src={userState.img} alt="Profile Image" className="user-img blhe" />
				<p className="user-detail">User Name: {userState.userName} </p>
				{!isVisibleName && (
					<input
						type="button"
						value="Change User Name"
						onClick={() => {
							setIsVisibleName(true);
						}}
						className="user-button"
					/>
				)}
				{isVisibleName && (
					<div>
						<input
							type="text"
							onChange={(e) => {
								setThisName(e.target.value);
							}}
							className="user-field"
							placeholder={userState.userName}
						/>
						<br />
						<br />
						<input type="button" value="Change User Name" onClick={handleClick} className="user-button" />
					</div>
				)}
				<p className="user-detail">Roll Number: {userState.rollNumber} </p>
			</div>
		</div>
	);
}

export default StudentProfile;
