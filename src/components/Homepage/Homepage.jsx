import aboutUs from '../../assets/text/about_us.json';
import "./Homepage.css"

function Homepage() {
    return (
        <div className="about-us">
            <h3> {aboutUs.heading} </h3>
            <div>{aboutUs.data}</div>
        </div>
    );
}

export default Homepage;