import React, { useState } from "react";
import "./Navbar.css";
import { Link } from "react-router-dom";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import { AppBar } from "@mui/material";

const Navbar = () => {
	return (
		<nav className="navbar">
			<div className="navbar-container">
				<div>
					<Link className="title" to={"dashboard-st"}>
						National Institute of Technology Hamirpur
					</Link>
				</div>
				<div className="navbar-profile">
					<Link>
						<AccountCircleIcon className="account-icon" fontSize="large" />
					</Link>
					<p className="account-name">UserName</p>
				</div>
			</div>
		</nav>
	);
};

export default Navbar;
