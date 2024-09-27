import Icon from "@mdi/react";
import {
  mdiHome,
  mdiFaceManProfile,
  mdiMessage,
  mdiHistory,
  mdiFileTree,
  mdiAccountGroup,
  mdiCog,
  mdiFaceAgent,
  mdiSecurity,
  mdiReact,
  mdiMagnify,
  mdiBellOutline,
} from "@mdi/js";

const App = () => {
  return (
    <div className="main-container">
      <nav className="sidebar-container sidebar">
        <a href="" className="sidebar-container-item menu-logo-anchor">
          <Icon className="menu-icon" path={mdiReact} size={2} color="white" />
          <span className="logo-button-label">Dashboard</span>
        </a>
        <div className="sidebar-menu-container sidebar-container-item">
          <ul className="main-menu">
            <li>
              <a className="menu-label-anchor" href="">
                <Icon
                  className="menu-icon"
                  path={mdiHome}
                  size={1}
                  color="white"
                />
                <span className="menu-button-label">Home</span>
              </a>
            </li>
            <li>
              <a className="menu-label-anchor" href="">
                <Icon
                  className="menu-icon"
                  path={mdiFaceManProfile}
                  size={1}
                  color="white"
                />
                <span className="menu-button-label">Profile</span>
              </a>
            </li>
            <li>
              <a className="menu-label-anchor" href="">
                <Icon
                  className="menu-icon"
                  path={mdiMessage}
                  size={1}
                  color="white"
                />
                <span className="menu-button-label">Messages</span>
              </a>
            </li>
            <li>
              <a className="menu-label-anchor" href="">
                <Icon
                  className="menu-icon"
                  path={mdiHistory}
                  size={1}
                  color="white"
                />
                <span className="menu-button-label">History</span>
              </a>
            </li>
            <li>
              <a className="menu-label-anchor" href="">
                <Icon
                  className="menu-icon"
                  path={mdiFileTree}
                  size={1}
                  color="white"
                />
                <span className="menu-button-label">Tasks</span>
              </a>
            </li>
            <li>
              <a className="menu-label-anchor" href="">
                <Icon
                  className="menu-icon"
                  path={mdiAccountGroup}
                  size={1}
                  color="white"
                />
                <span className="menu-button-label">Communities</span>
              </a>
            </li>
          </ul>
          <ul className="secondary-menu">
            <li>
              <a className="menu-label-anchor" href="">
                <Icon
                  className="menu-icon"
                  path={mdiCog}
                  size={1}
                  color="white"
                />
                <span className="menu-button-label">Settings</span>
              </a>
            </li>
            <li>
              <a className="menu-label-anchor" href="">
                <Icon
                  className="menu-icon"
                  path={mdiFaceAgent}
                  size={1}
                  color="white"
                />
                <span className="menu-button-label">Support</span>
              </a>
            </li>
            <li>
              <a className="menu-label-anchor" href="">
                <Icon
                  className="menu-icon"
                  path={mdiSecurity}
                  size={1}
                  color="white"
                />
                <span className="menu-button-label">Privacy</span>
              </a>
            </li>
          </ul>
        </div>
      </nav>
      <div className="body-container body">
        <div className="header-container header">
          <div className="search-container">
            <Icon className="search-icon" path={mdiMagnify} size={1} />
            <input className="searchbar" type="search" />
          </div>
          <div className="profile-container">
            <Icon
              className="profile-item-sm"
              path={mdiBellOutline}
              size={0.8}
            />
            <img className="profile-picture-md" src="/profile.jpg" />
            <div className="profile-sm">Morgan Oakley</div>
          </div>
          <div className="menu-container">
            <img className="profile-picture-lg" src="/profile.jpg" />
            <div className="hello-menu">
              <div className="salutation">Hi there,</div>
              <div className="profile-lg">Morgan Oakley (@morgan)</div>
            </div>
          </div>
          <div className="buttons-container">
            <button className="button-item">New</button>
            <button className="button-item">Upload</button>
            <button className="button-item">Share</button>
          </div>
        </div>
        <div className="maininfo-container maininfo">
          <div className="cards-container cards">
            <div className="cards-title">Your Projects</div>
            <div className="maincards-container">
              <div className="card card-container">
                <div className="card-title">Less Cool Project</div>
                <div className="card-body">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit.
                  Maiores, adipisci amet. Architecto laborum autem, accusantium
                  tempore impedit similique
                </div>
              </div>
              <div className="card card-container">
                <div className="card-title">Super Cool Project</div>
                <div className="card-body">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit.
                  Maiores, adipisci amet. Architecto laborum autem, accusantium
                  tempore impedit similique
                </div>
              </div>
              <div className="card card-container">
                <div className="card-title">Impossible App</div>
                <div className="card-body">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit.
                  Maiores, adipisci amet. Architecto laborum autem, accusantium
                  tempore impedit similique
                </div>
              </div>
              <div className="card card-container">
                <div className="card-title">Easy Peasy App</div>
                <div className="card-body">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit.
                  Maiores, adipisci amet. Architecto laborum autem, accusantium
                  tempore impedit similique
                </div>
              </div>
              <div className="card card-container">
                <div className="card-title">Easy Peasy App</div>
                <div className="card-body">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit.
                  Maiores, adipisci amet. Architecto laborum autem, accusantium
                  tempore impedit similique
                </div>
              </div>
              <div className="card card-container">
                <div className="card-title">Money Maker</div>
                <div className="card-body">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit.
                  Maiores, adipisci amet. Architecto laborum autem, accusantium
                  tempore impedit similique
                </div>
              </div>
            </div>
          </div>
          <div className="extras-container extras">
            <div className="extras-title">Announcements</div>
            <div className="extra-container">
              <div className="announcements-item">
                <div className="announcements-item-title">Site Maintenance</div>
                <div className="announcements-item-body">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit. Lorem
                  ipsum dolor sit. Lorem ipsum dolor, sit amet
                </div>
              </div>
              <div className="announcements-item">
                <div className="announcements-item-title">
                  Community Share Day
                </div>
                <div className="announcements-item-body">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit. Lorem
                  ipsum dolor sit. Lorem ipsum dolor, sit amet
                </div>
              </div>
              <div className="announcements-item">
                <div className="announcements-item-title">
                  Updated Privacy Policy
                </div>
                <div className="announcements-item-body">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit. Lorem
                  ipsum dolor sit. Lorem ipsum dolor, sit amet
                </div>
              </div>
            </div>
            <div className="extras-title">Trending</div>
            <div className="extra-container">
              <div className="trending-item">
                <img className="trending-item-pfp" src="/profile.jpg" />
                <div className="trending-item-body">
                  <div className="trending-item-body-username">@tegan</div>
                  <div className="trending-item-body-description">
                    World Peace Builder
                  </div>
                </div>
              </div>
              <div className="trending-item">
                <img className="trending-item-pfp" src="/profile.jpg" />
                <div className="trending-item-body">
                  <div className="trending-item-body-username">@morgan</div>
                  <div className="trending-item-body-description">
                    Super Cool Project
                  </div>
                </div>
              </div>
              <div className="trending-item">
                <img className="trending-item-pfp" src="/profile.jpg" />
                <div className="trending-item-body">
                  <div className="trending-item-body-username">@kendall</div>
                  <div className="trending-item-body-description">
                    Life Changing App
                  </div>
                </div>
              </div>
              <div className="trending-item">
                <img className="trending-item-pfp" src="/profile.jpg" />
                <div className="trending-item-body">
                  <div className="trending-item-body-username">alex</div>
                  <div className="trending-item-body-description">
                    No Traffic Maker
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default App;
