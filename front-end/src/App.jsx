import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./Components/Navbar/Navbar";
import StudentDashboard from "./Pages/Student Dashboard/StudentDashboard";
import Sidebar from "./Components/Sidebar/Sidebar";
import StudentProfile from "./Pages/Student Profile/StudentProfile";
import { useState, useEffect, createContext } from "react";
import Setup from "./Pages/Login/Setup";
import TeacherDashboard from "./Pages/Teacher Dashboard/TeacherDashboard";

export const AppContext = createContext();

function App() {
  const [accesstoken, setAccesstoken] = useState(null);
  const [userState, setUserState] = useState(
    JSON.parse(localStorage.getItem("userState")) || {
      email: "",
      firstname: "",
      lastname: "",
      isfaculty: false,
      isloggedin: false,
      level: "",
    }
  );

  useEffect(() => {
    localStorage.setItem("userState", JSON.stringify(userState));
  }, [userState]);

  return (
    <div className="App">
      <AppContext.Provider
        value={{ userState, setUserState, setAccesstoken, accesstoken }}
      >
        <Router>
          <Navbar />
          <Sidebar />
          <Routes>
            <Route element={<Setup />} exact path="/" />
            <Route element={<StudentDashboard />} exact path="/dashboard-st" />
            <Route
              element={<TeacherDashboard />}
              exact
              path="/dashboard-ty"
            />
            <Route element={<StudentProfile />} exact path="/profile-st" />
          </Routes>
        </Router>
      </AppContext.Provider>
    </div>
  );
}

export default App;
