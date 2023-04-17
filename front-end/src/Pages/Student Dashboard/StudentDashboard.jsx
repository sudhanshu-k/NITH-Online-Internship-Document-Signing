import React, { useEffect } from "react";
import "./StudentDashboard.css";
import { useContext } from "react";
import { AppContext } from "../../App";
import axios from "axios";

function StudentDashboard() {
	const { accesstoken } = useContext(AppContext);
	let header = {
		Authorization: `Bearer ${accesstoken}`,
	};

	useEffect(() => {
		axios.get("http://127.0.0.1:3000/api/profile/dashboard", { headers: header }).then((response) => {
			console.log(response);
		});
	}, []);

	return <div className="dashboard-body">StudentDashboard</div>;
}

export default StudentDashboard;
