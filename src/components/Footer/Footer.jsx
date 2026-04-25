import React from "react";
import './Footer.css'
function Footer() {
    return (
        <footer className="footer">
            <div className="footer-container">
                <div className="footer-links">

                    <ul>
                        <li> <a href="#">GitHub</a></li>
                        <li> <a href="#">LinkedIn</a></li>
                        <li> <a href="#">YouTube</a></li>
                    </ul>

                </div>

                <div className="footer-bottom">
                    <p> &copy; {new Date().getFullYear()} Mindfull Education. All rights reserved.</p>
                </div>

            </div>
        </footer>
    )
}

export default Footer;