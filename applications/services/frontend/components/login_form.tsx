import { useSession, signIn, signOut } from "next-auth/react";

const LoginForm = () => {
    const { data: session } = useSession()
    return (
        <div className="row g-3 align-items-center">
        {!session ? (
            <>
                <button className="btn btn-primary" onClick={() => signIn("github")}>
                    Sign in with Github <img width="25" src="https://cdn-icons-png.flaticon.com/512/25/25231.png" alt="GH"/>
                </button>
            </>
        ) : (
            <>
                <p className="text-center">
                    Not {session.user && (session.user.name || session.user.email)}? Then
                    Logout and login again
                </p>
                <button className="btn btn-primary" onClick={() => signOut()}>Logout</button> <br />
            </>
        )}
        </div>
    )
}
export default LoginForm