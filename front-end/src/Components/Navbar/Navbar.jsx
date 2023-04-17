import React, { useState } from "react";
import "./Navbar.css";
import { Link } from "react-router-dom";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import { useContext } from "react";
import { AppContext } from "../../App";

const Navbar = () => {
	const { userState } = useContext(AppContext);
	// console.log(userState);
	return (
		<nav className="navbar">
			<div className="navbar-container">
				<div>
					<Link className="title" to={"dashboard-st"}>
						National Institute of Technology Hamirpur
					</Link>
				</div>
				<div className="navbar-profile">
					<Link className="profile-container" to={"/profile-st"}>
						<AccountCircleIcon className="account-icon" fontSize="large" />
						<p className="account-name">{userState.firstname}</p>
					</Link>
				</div>
			</div>
		</nav>
	);
};

export default Navbar;
