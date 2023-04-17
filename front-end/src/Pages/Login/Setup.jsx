import React from "react";
import { useState } from "react";
import Signin from "./Signin";
import Login from "./Login";
import "./Login.css";

function Setup() {
	const [flag, setFlag] = useState(false);
	return (
		<div className="login-body setup-body">
			<div className="setup-child">{flag ? <Login /> : <Signin func={setFlag} />}</div>
			<div className="setup-child">
				<button className="flag-btn login-submit" onClick={() => setFlag(!flag)}>
					{flag ? " Create Account..." : "Already have an account"}
				</button>
			</div>
		</div>
	);
}

export default Setup;
