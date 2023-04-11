import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./Components/Navbar/Navbar";
import StudentDashboard from "./Pages/Student Dashboard/StudentDashboard";
import Sidebar from "./Components/Sidebar/Sidebar";
import StudentProfile from "./Pages/Student Profile/StudentProfile";
import Login from "./Pages/Login/Login";
function App() {
	return (
		<div className="App">
			<Router>
				<Navbar />
				<Sidebar />
				<Routes>
					<Route element={<Login />} exact path="/login" />
					<Route element={<StudentDashboard />} exact path="/dashboard-st" />
					<Route element={<StudentProfile />} exact path="/profile-st" />
				</Routes>
			</Router>
		</div>
	);
}

export default App;
