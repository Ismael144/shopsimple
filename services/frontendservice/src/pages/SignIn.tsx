import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Link } from 'react-router';

export default function SignIn() {
    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        const data = new FormData(event.currentTarget)

        // Make api request
    }

    return (
        <div className="container d-flex align-center justify-content-between h-100" style={{ height: "100vh" }}>
            <div className="auth-form bg-white m-auto p-4 mt-5">
                <form action="" method="post" onSubmit={handleSubmit}>
                    <h2 className="text-secondary text-center my-3">Sign In</h2>
                    <div className="form-element my-3">
                        <label htmlFor="" className="form-label">Username: </label>
                        <input type="username" name="username" className="form-control form-control-md" placeholder="Enter your username..." />
                    </div>
                    <div className="form-element my-3">
                        <label htmlFor="" className="form-label">Password: </label>
                        <input type="password" name="password" className="form-control form-control-md" placeholder="Enter your password..." />
                    </div>
                    <div className="d-flex align-items-center gap-2">
                        <input type="checkbox" name="" id="" />
                        <label htmlFor="" className="form-label m-0 p-0">Remember Me</label>
                    </div>
                    <input type="submit" value="Sign In" className="my-3 btn btn-lg btn-success w-100" />
                    <div className="text-center mb-3">
                        Don't have an account? <Link to="/signup">Sign Up</Link>
                    </div>
                </form>
            </div>
        </div>
    );
}