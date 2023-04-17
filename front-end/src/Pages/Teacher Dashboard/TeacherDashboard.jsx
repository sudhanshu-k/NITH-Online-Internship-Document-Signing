import React from "react";
import UgIntern from "../Forms/UG Training Request Form/UgInten";

import styles from "./TeacherDashboard.module.css";


function TeacherDashboard() {

	const data={
		"name" : "Kshitij Roodkee",
		"fname": "ASDH",
		"rollNo" : "20bcs025",
		"department": "CSE",
		"programme":"Btech",
		"email" :"kshitijroodkee1@gmail.com",
		"programme":"B.Tech",
		"phone": "8091734849"

	}


	return <>
<div>TeacherDashboard</div>
<div className={styles.control}>
	<button className={styles.approve}> Approve </button>
	<button className={styles.reject}> Reject</button>
</div>
<UgIntern {...data}/>
	</> 
}

export default TeacherDashboard;
