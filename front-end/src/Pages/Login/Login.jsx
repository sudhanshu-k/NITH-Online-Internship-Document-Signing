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

function Login() {
	const { userState, setUserState, setAccesstoken } = useContext(AppContext);

	let schema = yup.object().shape({
		email: yup.string().email().required(),
		password: yup.string().min(8).required(),
	});

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({ resolver: yupResolver(schema) });

	const navigate = useNavigate();

	const onSubmit = (data) => {
		axios
			.post("http://127.0.0.1:3000/api/auth/signin", data, {
				withCredentials: true,
				credential: "include",
			})
			.then(function (response) {
				if (response.status == 200) {
					console.log(response.data.user);
					// setAccesstoken(response.data.data.access_token);
					setUserState(response.data.user);
					navigate("/dashboard-st");
					// console.log(userState);
				} else {
					alert("Something Went Wrong");
				}
			})
			.catch(function (error) {
				console.log(error);
			});
	};

	return (
		<StylesProvider>
			<div className="login-card">
				<div className="login-header"></div>
				<div className="login-form">
					<form onSubmit={handleSubmit(onSubmit)}>
						<input type="text" className="login-input" placeholder="Email..." {...register("email")} />
						<p className="login-error">{errors.email?.message}</p>
						<input type="password" className="login-input" placeholder="Password..." {...register("password")} />
						<p className="login-error">{errors.password?.message}</p>
						<input type="submit" className="login-submit" />
					</form>
				</div>
			</div>
		</StylesProvider>
	);
}

export default Login;
