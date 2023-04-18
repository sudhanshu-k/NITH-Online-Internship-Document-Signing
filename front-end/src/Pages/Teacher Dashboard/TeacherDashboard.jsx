import React, { useEffect, useState } from "react";
import UgIntern from "../Forms/UG Training Request Form/UgInten";
import styles from "./TeacherDashboard.module.css";
import axios from "axios";

function TeacherDashboard() {
	const [dataForms, setDataForms] = useState({});

	useEffect(() => {
		async function fetchDashboardData() {
			try {
				const response = await axios
					.get("http://127.0.0.1:3000/api/profile/dashboard", {
						withCredentials: true,
						credentials: "include",
					})
					.then((response) => {
						console.log(response.data.rows);
						setDataForms(response.data.rows);
					});
				return response.data.rows; // return the data
			} catch (error) {
				console.error(error);
				throw error; // throw error to be caught by the calling function
			}
		}

		async function setData() {
			const data = await fetchDashboardData();
			setDataForms(data);
			console.log("Static Data");
			console.log(dataForms);
		}
		// fetchDashboardData();
		setData(); // call the async function
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
					<div className={styles.formTab} key={form.user.roll}>
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
