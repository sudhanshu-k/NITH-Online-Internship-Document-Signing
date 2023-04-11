import React from "react";
import { Card, CardContent } from "@mui/material";
import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import * as yup from "yup";
import "./Login.css";
import { StylesProvider } from "@mui/styles";
import axios from "axios";

function Login() {
	let schema = yup.object().shape({
		email: yup.string().email().required(),
		password: yup.string().min(8).required(),
	});

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({ resolver: yupResolver(schema) });

	const onSubmit = (data) => {
		axios
			.post("/api/auth/signin", data)
			.then(function (response) {
				console.log(response);
			})
			.catch(function (error) {
				console.log(error);
			});
	};

	return (
		<StylesProvider>
			<div className="login-body">
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
			</div>
		</StylesProvider>
	);
}

export default Login;
