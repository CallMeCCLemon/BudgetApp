"use client";
import {useState} from "react";
import "@/components/forms/forms.css";

export default function Page() {
    const [inputs, setInputs] = useState({
        username: "",
        password: "",
    });

    const handleChange = (event: { target: { name: any; value: any; }; }) => {
        const name = event.target.name;
        const value = event.target.value;
        setInputs(values => ({...values, [name]: value}))
    }

    const handleSubmit = (event: { preventDefault: () => void; }) => {
        event.preventDefault();
        alert(inputs);
    }

    return (
        <div className="wrapper">
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Enter your username or email:
                        <input
                            type="text"
                            name="username"
                            value={inputs.username || ""}
                            onChange={handleChange}
                        />
                    </label>
                </div>
                <div>
                    <label>Enter your password:
                        <input
                            type="password"
                            name="password"
                            value={inputs.password || ""}
                            onChange={handleChange}
                        />
                    </label>
                </div>
                <div>
                    <input type="submit" />
                </div>
            </form>
        </div>
    )
}
