import React, { useEffect, useState } from "react";
import { X } from "react-feather";

import InputControl from "../InputControl/InputControl";

import styles from "./Editor.module.css";
import axios from "axios";

function Editor(props) {
	const sections = props.sections;
	const information = props.information;

	const [activeSectionKey, setActiveSectionKey] = useState(Object.keys(sections)[0]);
	const [activeInformation, setActiveInformation] = useState(information[sections[Object.keys(sections)[0]]]);
	const [activeDetailIndex, setActiveDetailIndex] = useState(0);
	const [sectionTitle, setSectionTitle] = useState(sections[Object.keys(sections)[0]]);
	const [values, setValues] = useState({
		name: activeInformation?.detail?.name || "",
		rollNo: activeInformation?.detail?.rollNo || "",
		fName: activeInformation?.detail?.fName || "",
		programme: activeInformation?.detail?.programme || "",
		department: activeInformation?.detail?.department || "cse",
		phone: activeInformation?.detail?.phone || "",
		programme: activeInformation?.detail?.programme || "",
		rollNo: activeInformation?.detail?.rollNo || "",
		fName: activeInformation?.detail?.fName || "",
		email: activeInformation?.detail?.email || "",
	});

	const handlePointUpdate = (value, index) => {
		const tempValues = { ...values };
		if (!Array.isArray(tempValues.points)) tempValues.points = [];
		tempValues.points[index] = value;
		setValues(tempValues);
	};

	const basicInfoBody = (
		<div className={styles.detail}>
			<div className={styles.row}>
				<InputControl
					label="Name"
					placeholder="Enter your full name eg. Aashu"
					value={values.name}
					onChange={(event) => setValues((prev) => ({ ...prev, name: event.target.value }))}
				/>
				<InputControl
					label="Roll No"
					value={values.rollNo}
					placeholder="Enter your Roll no eg. 20bcs025"
					onChange={(event) => setValues((prev) => ({ ...prev, rollNo: event.target.value }))}
				/>
			</div>
			<div className={styles.row}>
				<InputControl
					label="Father's Name"
					value={values.fName}
					placeholder="Enter your Father's Name"
					onChange={(event) => setValues((prev) => ({ ...prev, fName: event.target.value }))}
				/>
				<InputControl
					label="Programme"
					value={values.programme}
					placeholder="Eg. B.Tech"
					onChange={(event) => setValues((prev) => ({ ...prev, programme: event.target.value }))}
				/>
			</div>
			<InputControl
				label="Department"
				value={values.department}
				placeholder="Eg. CSE"
				onChange={(event) => setValues((prev) => ({ ...prev, department: event.target.value }))}
			/>
			<div className={styles.row}>
				<InputControl
					label="Email"
					value={values.email}
					placeholder="Enter your email"
					onChange={(event) => setValues((prev) => ({ ...prev, email: event.target.value }))}
				/>
				<InputControl
					label="Enter phone"
					value={values.phone}
					placeholder="Enter your phone number"
					onChange={(event) => setValues((prev) => ({ ...prev, phone: event.target.value }))}
				/>
			</div>
		</div>
	);

	const generateBody = () => {
		switch (sections[activeSectionKey]) {
			case sections.basicInfo:
				return basicInfoBody;
			case sections.workExp:
				return workExpBody;
			case sections.project:
				return projectBody;
			case sections.education:
				return educationBody;
			case sections.achievement:
				return achievementsBody;
			case sections.summary:
				return summaryBody;
			case sections.other:
				return otherBody;
			default:
				return null;
		}
	};

	const handleSubmission = () => {
		switch (sections[activeSectionKey]) {
			case sections.basicInfo: {
				const tempDetail = {
					name: values.name,
					rollNo: values.rollNo,
					programme: values.programme,
					department: values.department,
					email: values.email,
					phone: values.phone,
					fName: values.fName,
				};

				props.setInformation((prev) => ({
					...prev,
					[sections.basicInfo]: {
						...prev[sections.basicInfo],
						detail: tempDetail,
						sectionTitle,
					},
				}));
				break;
			}
			case sections.workExp: {
				const tempDetail = {
					certificationLink: values.certificationLink,
					title: values.title,
					startDate: values.startDate,
					endDate: values.endDate,
					companyName: values.companyName,
					location: values.location,
					points: values.points,
				};
				const tempDetails = [...information[sections.workExp]?.details];
				tempDetails[activeDetailIndex] = tempDetail;

				props.setInformation((prev) => ({
					...prev,
					[sections.workExp]: {
						...prev[sections.workExp],
						details: tempDetails,
						sectionTitle,
					},
				}));
				break;
			}
			case sections.project: {
				const tempDetail = {
					link: values.link,
					title: values.title,
					overview: values.overview,
					github: values.github,
					points: values.points,
				};
				const tempDetails = [...information[sections.project]?.details];
				tempDetails[activeDetailIndex] = tempDetail;

				props.setInformation((prev) => ({
					...prev,
					[sections.project]: {
						...prev[sections.project],
						details: tempDetails,
						sectionTitle,
					},
				}));
				break;
			}
			case sections.education: {
				const tempDetail = {
					title: values.title,
					college: values.college,
					startDate: values.startDate,
					endDate: values.endDate,
				};
				const tempDetails = [...information[sections.education]?.details];
				tempDetails[activeDetailIndex] = tempDetail;

				props.setInformation((prev) => ({
					...prev,
					[sections.education]: {
						...prev[sections.education],
						details: tempDetails,
						sectionTitle,
					},
				}));
				break;
			}
			case sections.achievement: {
				const tempPoints = values.points;

				props.setInformation((prev) => ({
					...prev,
					[sections.achievement]: {
						...prev[sections.achievement],
						points: tempPoints,
						sectionTitle,
					},
				}));
				break;
			}
			case sections.summary: {
				const tempDetail = values.summary;

				props.setInformation((prev) => ({
					...prev,
					[sections.summary]: {
						...prev[sections.summary],
						detail: tempDetail,
						sectionTitle,
					},
				}));
				break;
			}
			case sections.other: {
				const tempDetail = values.other;

				props.setInformation((prev) => ({
					...prev,
					[sections.other]: {
						...prev[sections.other],
						detail: tempDetail,
						sectionTitle,
					},
				}));
				break;
			}
		}
	};

	const postReq = (data) => {
		console.log(data);
		axios
			.post("http://127.0.0.1:3000/api/form/ugintern", data, {
				withCredentials: true,
				credential: "include",
			})
			.then(function (response) {
				if (response.status == 200) {
					console.log(response);
				} else {
					alert("Something Went Wrong");
				}
			})
			.catch(function (error) {
				console.log(error);
			});
	};

	const handleForms = () => {
		switch (sections[activeSectionKey]) {
			case sections.basicInfo: {
				const tempDetail = {
					name: values.name,
					fathername: values.fName,
					address: "address",
					contact: values.phone,
					companyname: "companynamed",
					email: values.email,
					aoi: "aoi",
					isoffline: true,
					startday: "2023-04-15T00:00:00Z",
					endday: "2023-04-30T00:00:00Z",
					weeks: 2,
					formtpo: true,
					stipend: 1000,
					fromdate: "2023-04-15T00:00:00Z",
					remarksdept: "Some dept remarks",
					remarksfi: "Some FI remarks",
				};
				postReq(tempDetail);
				break;
			}
		}
	};

	const handleAddNew = () => {
		const details = activeInformation?.details;
		if (!details) return;
		const lastDetail = details.slice(-1)[0];
		if (!Object.keys(lastDetail).length) return;
		details?.push({});

		props.setInformation((prev) => ({
			...prev,
			[sections[activeSectionKey]]: {
				...information[sections[activeSectionKey]],
				details: details,
			},
		}));
		setActiveDetailIndex(details?.length - 1);
	};

	const handleDeleteDetail = (index) => {
		const details = activeInformation?.details ? [...activeInformation?.details] : "";
		if (!details) return;
		details.splice(index, 1);
		props.setInformation((prev) => ({
			...prev,
			[sections[activeSectionKey]]: {
				...information[sections[activeSectionKey]],
				details: details,
			},
		}));

		setActiveDetailIndex((prev) => (prev === index ? 0 : prev - 1));
	};

	useEffect(() => {
		const activeInfo = information[sections[activeSectionKey]];
		setActiveInformation(activeInfo);
		setSectionTitle(sections[activeSectionKey]);
		setActiveDetailIndex(0);
		setValues({
			name: activeInfo?.detail?.name || "",
			overview: activeInfo?.details ? activeInfo.details[0]?.overview || "" : "",
			link: activeInfo?.details ? activeInfo.details[0]?.link || "" : "",
			certificationLink: activeInfo?.details ? activeInfo.details[0]?.certificationLink || "" : "",
			companyName: activeInfo?.details ? activeInfo.details[0]?.companyName || "" : "",
			college: activeInfo?.details ? activeInfo.details[0]?.college || "" : "",
			location: activeInfo?.details ? activeInfo.details[0]?.location || "" : "",
			startDate: activeInfo?.details ? activeInfo.details[0]?.startDate || "" : "",
			endDate: activeInfo?.details ? activeInfo.details[0]?.endDate || "" : "",
			points: activeInfo?.details
				? activeInfo.details[0]?.points
					? [...activeInfo.details[0]?.points]
					: ""
				: activeInfo?.points
				? [...activeInfo.points]
				: "",
			title: activeInfo?.details ? activeInfo.details[0]?.title || "" : activeInfo?.detail?.title || "",
			linkedin: activeInfo?.detail?.linkedin || "",
			github: activeInfo?.details ? activeInfo.details[0]?.github || "" : activeInfo?.detail?.github || "",
			phone: activeInfo?.detail?.phone || "",
			department: activeInfo?.detail?.department || "cse",
			programme: activeInfo?.detail?.programme || "",
			rollNo: activeInfo?.detail?.rollNo || "",
			fName: activeInfo?.detail?.fName || "",
			email: activeInfo?.detail?.email || "",
			summary: typeof activeInfo?.detail !== "object" ? activeInfo.detail : "",
			other: typeof activeInfo?.detail !== "object" ? activeInfo.detail : "",
		});
	}, [activeSectionKey]);

	useEffect(() => {
		setActiveInformation(information[sections[activeSectionKey]]);
	}, [information]);

	useEffect(() => {
		const details = activeInformation?.details;
		if (!details) return;

		const activeInfo = information[sections[activeSectionKey]];
		setValues({
			overview: activeInfo.details[activeDetailIndex]?.overview || "",
			link: activeInfo.details[activeDetailIndex]?.link || "",
			certificationLink: activeInfo.details[activeDetailIndex]?.certificationLink || "",
			companyName: activeInfo.details[activeDetailIndex]?.companyName || "",
			location: activeInfo.details[activeDetailIndex]?.location || "",
			startDate: activeInfo.details[activeDetailIndex]?.startDate || "",
			endDate: activeInfo.details[activeDetailIndex]?.endDate || "",
			points: activeInfo.details[activeDetailIndex]?.points || "",
			title: activeInfo.details[activeDetailIndex]?.title || "",
			linkedin: activeInfo.details[activeDetailIndex]?.linkedin || "",
			github: activeInfo.details[activeDetailIndex]?.github || "",
			college: activeInfo.details[activeDetailIndex]?.college || "",
		});
	}, [activeDetailIndex]);

	return (
		<div className={styles.container}>
			<div className={styles.header}>
				{Object.keys(sections)?.map((key) => (
					<div
						className={`${styles.section} ${activeSectionKey === key ? styles.active : ""}`}
						key={key}
						onClick={() => setActiveSectionKey(key)}
					>
						{sections[key]}
					</div>
				))}
			</div>

			<div className={styles.body}>
				<InputControl
					label="Title"
					placeholder="Enter section title"
					value={sectionTitle}
					onChange={(event) => setSectionTitle(event.target.value)}
				/>

				<div className={styles.chips}>
					{activeInformation?.details
						? activeInformation?.details?.map((item, index) => (
								<div
									className={`${styles.chip} ${activeDetailIndex === index ? styles.active : ""}`}
									key={item.title + index}
									onClick={() => setActiveDetailIndex(index)}
								>
									<p>
										{sections[activeSectionKey]} {index + 1}
									</p>
									<X
										onClick={(event) => {
											event.stopPropagation();
											handleDeleteDetail(index);
										}}
									/>
								</div>
						  ))
						: ""}
					{activeInformation?.details && activeInformation?.details?.length > 0 ? (
						<div className={styles.new} onClick={handleAddNew}>
							+New
						</div>
					) : (
						""
					)}
				</div>

				{generateBody()}

				<button onClick={handleSubmission}>Save</button>
			</div>
			<div>
				<button className={styles.submit} onClick={handleForms}>
					Submit
				</button>
			</div>
		</div>
	);
}

export default Editor;
