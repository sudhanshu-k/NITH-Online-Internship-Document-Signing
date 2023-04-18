import Body from "../../Components/Body/Body";
import React, { useEffect } from "react";
import "./StudentDashboard.css";
import { useContext } from "react";
import { AppContext } from "../../App";
import axios from "axios";

function StudentDashboard() {
	const { accesstoken } = useContext(AppContext);
	useEffect(() => {
		axios.get("http://127.0.0.1:3000/api/profile/dashboard").then(function (response) {
			console.log(response);
		});
	}, []);

	return (
		<>
			<div className="dashboard-body">StudentDashboard</div>
			<div className="form">
				<Body />
			</div>
		</>
	);
}

export default StudentDashboard;
