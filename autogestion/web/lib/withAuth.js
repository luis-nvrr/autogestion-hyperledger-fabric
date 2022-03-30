import Index from "../pages/index";
import useUser from "../lib/useUser";

const withAuth = (Component) => {
  const Auth = (props) => {
    const { isLoggedIn } = useUser();

    if (!isLoggedIn) {
      return <Index />;
    }

    return <Component {...props} />;
  };

  if (Component.getInitialProps) {
    Auth.getInitialProps = Component.getInitialProps;
  }

  return Auth;
};

export default withAuth;
