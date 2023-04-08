import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./Components/Navbar/Navbar";
import StudentDashboard from "./Pages/Student Dashboard/StudentDashboard";
import Sidebar from "./Components/Sidebar/Sidebar";
function App() {
	return (
		<div className="App">
			<Router>
				<Navbar />
				<Sidebar />
				<Routes>
					<Route elemment={<StudentDashboard />} path="dashboard-st" />
				</Routes>
			</Router>
		</div>
	);
}

export default App;
