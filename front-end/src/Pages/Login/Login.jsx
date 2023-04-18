import React from "react";
import { useContext, useState } from "react";
import { AppContext } from "../../App";
import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import * as yup from "yup";
import "./Login.css";
import { StylesProvider } from "@mui/styles";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { ToggleButton } from "@mui/material";
import { Check } from "@mui/icons-material";

function Login() {
	const { userState, setUserState, setAccesstoken } = useContext(AppContext);

	let schema = yup.object().shape({
		email: yup.string().email().required(),
		password: yup.string().min(5).required(),
	});

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({ resolver: yupResolver(schema) });

	const navigate = useNavigate();
	const [selected, setSelected] = React.useState(false);

	const onSubmit = (data) => {
		console.log("Submitted");
		var responseGet = {};
		axios
			.post("http://127.0.0.1:3000/api/auth/signin", data, {
				withCredentials: true,
				credential: "include",
			})
			.then(function (response) {
				responseGet = response;
				// console.log(responseGet.data);
				setUserState(responseGet.data.user);
				if (responseGet.status == 200) {
					setUserState(responseGet.data.user);
					// console.log(responseGet.data.user.isfaculty);
					if (responseGet.data.user.isfaculty == true) {
						navigate("/dashboard-ty");
					} else {
						navigate("/dashboard-st");
					}
				}
			})
			.catch(function (error) {
				console.log(error);
			});
	};

	return (
		<StylesProvider>
			<div className="login-card">
				<div className="login-form">
					<form onSubmit={handleSubmit(onSubmit)}>
						<input type="text" className="login-input" placeholder="Email..." {...register("email")} />
						<p className="login-error">{errors.email?.message}</p>
						<input type="password" className="login-input" placeholder="Password..." {...register("password")} />
						<p className="login-error">{errors.password?.message}</p>
						<input type="submit" className="login-submit" />
					</form>
				</div>
				<div className="login-teach">Are you a faculty member ?</div>
				<ToggleButton
					value="check"
					selected={selected}
					color="success"
					size="small"
					onChange={() => {
						setSelected(!selected);
					}}
				>
					<Check />
				</ToggleButton>
			</div>
		</StylesProvider>
	);
}

export default Login;
