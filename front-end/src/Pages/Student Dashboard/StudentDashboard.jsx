

import UgIntern from "../Forms/UG Training Request Form/UgInten";

import Body from "../../Components/Body/Body";



let formProps = {
name: "Kshitij Roodkee",
fname:"Jitender Singh Roodkee"

}




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

	return (<>
  <div className="dashboard-body">StudentDashboard</div>
  <div className="form"><Body/></div>
  </>);

}

export default StudentDashboard;
