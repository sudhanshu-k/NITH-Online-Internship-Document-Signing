import React, { useEffect } from "react";
import UgIntern from "../Forms/UG Training Request Form/UgInten";
import styles from "./TeacherDashboard.module.css";
import axios from "axios";

function TeacherDashboard() {
	useEffect(() => {
		axios.get("http://127.0.0.1:3000/api/profile/dashboard");
	}, []);

	const data = {
		name: "Kshitij Roodkee",
		fname: "ASDH",
		rollNo: "20bcs025",
		department: "CSE",
		programme: "Btech",
		email: "kshitijroodkee1@gmail.com",
		programme: "B.Tech",
		phone: "8091734849",
	};

	const forms = [
		{ uuid: "qwertyuiop", type: "UG Intern", user: { name: "Maharshi", roll: "20bcs004" } },
		{ uuid: "asdfghjkl", type: "UG Intern", user: { name: "Kshitij", roll: "20bcs025" } },
		{ uuid: "zxcvbnm", type: "UG Intern", user: { name: "Sudhanshu", roll: "20bcs083" } },
	];

	const handleClick = (uuid) => {
		console.log(uuid);
		// get request with input of uuid
	};

	return (
		<div className={styles.tyBody}>
			{forms.map((form) => {
				return (
					<div className={styles.formTab}>
						<div className={styles.formTabChild}> {form.uuid}</div>
						<div className={styles.formTabChild}> {form.type}</div>
						<div className={styles.formTabChild}> {form.user.name}</div>
						<div className={styles.formTabChild}> {form.user.roll}</div>
						<div className={styles.formTabChild}>
							<button
								className={styles.formTabButton}
								onClick={() => {
									handleClick(form.uuid);
								}}
							>
								Open
							</button>
						</div>
					</div>
				);
			})}
			{/* <div className={styles.control}>
				<button className={styles.approve}> Approve </button>
				<button className={styles.reject}> Reject</button>
			</div>
			<UgIntern {...data} /> */}
		</div>
	);
}

export default TeacherDashboard;
