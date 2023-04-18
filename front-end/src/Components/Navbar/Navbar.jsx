import React, { useState } from "react";
import "./Navbar.css";
import { Link } from "react-router-dom";
import { useContext } from "react";
import { AppContext } from "../../App";

const Navbar = () => {
	const { userState } = useContext(AppContext);
	return (
		<nav className="navbar">
			<div className="navbar-container">
				<div>
					<Link className="title" to={"dashboard-st"}>
						National Institute of Technology Hamirpur
					</Link>
				</div>
				<div className="navbar-profile">
					{userState.firstname != "" ? (
						<Link className="profile-container" to={"/profile-st"}>
							<p className="account-name">{userState.firstname + " " + userState.lastname}</p>
						</Link>
					) : (
						<Link className="profile-container" to={"/"}>
							<p className="account-name">Login</p>
						</Link>
					)}
					{/* <Link className="profile-container" to={"/profile-st"}>
						<p className="account-name">{userState.firstname}</p>
					</Link> */}
				</div>
			</div>
		</nav>
	);
};

export default Navbar;
