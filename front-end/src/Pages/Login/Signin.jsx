import React from "react";
import { useContext, useState } from "react";
import { AppContext } from "../../App";
import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import * as yup from "yup";
import { StylesProvider } from "@mui/styles";
import axios from "axios";
import { useNavigate } from "react-router-dom";

function Signin(props) {
	const { userState, setUserState } = useContext(AppContext);
	const [newState, setNewState] = useState({});

	let schema2 = yup.object().shape({
		FirstName: yup.string().required(),
		LastName: yup.string().required(),
		Email: yup.string().email().required(),
		Password: yup.string().min(8).required(),
	});

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({ resolver: yupResolver(schema2) });

	const navigate = useNavigate();

	const onSubmit = (data) => {
		axios.post("http://127.0.0.1:3000/api/auth/register", data).then(
			(response) => {
				console.log(response);
				if (response.status == 200) {
					props.func(true);
					navigate("/");
				} else {
					alert("Something Went Wrong");
				}
			},
			(error) => {
				console.log(error);
			},
		);
	};

	return (
		<StylesProvider>
			<div>
				<div className="login-card">
					<div className="login-header"></div>
					<div className="login-form">
						<form onSubmit={handleSubmit(onSubmit)}>
							<input type="text" className="login-input" placeholder="First Name..." {...register("FirstName")} />
							<p className="login-error">{errors.FirstName?.message}</p>
							<input type="text" className="login-input" placeholder="Last Name..." {...register("LastName")} />
							<p className="login-error">{errors.Email?.message}</p>
							<input type="text" className="login-input" placeholder="Email..." {...register("Email")} />
							<p className="login-error">{errors.Email?.message}</p>
							<input type="password" className="login-input" placeholder="Password..." {...register("Password")} />
							<p className="login-error">{errors.Password?.message}</p>
							<input type="submit" className="login-submit" />
						</form>
					</div>
				</div>
			</div>
		</StylesProvider>
	);
}

export default Signin;
