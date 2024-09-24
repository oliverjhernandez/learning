const App = () => {
  return (
    <div className="main-container">
      <nav className="sidebar-container sidebar">
        <div className="logo">Dashboard</div>
        <ul>
          <li>
            <a href="">Home</a>
          </li>
          <li>
            <a href="">Profile</a>
          </li>
          <li>
            <a href="">Messages</a>
          </li>
          <li>
            <a href="">History</a>
          </li>
          <li>
            <a href="">Tasks</a>
          </li>
          <li>
            <a href="">Communities</a>
          </li>
          <li>
            <a href="">Settings</a>
          </li>
          <li>
            <a href="">Support</a>
          </li>
          <li>
            <a href="">Privacy</a>
          </li>
        </ul>
      </nav>
      <div className="body-container body">
        <div className="header-container header">
          <div>SearchBar</div>
          <div>Profile</div>
          <div>Menu</div>
          <div>Buttons</div>
        </div>
        <div className="maininfo-container">
          <div className="cards-container cards">
            <div className="cards-title">Cards Title</div>
            <div className="maincards-container">
              <div className="card">Card1</div>
              <div className="card">Card2</div>
              <div className="card">Card3</div>
              <div className="card">Card4</div>
              <div className="card">Card5</div>
              <div className="card">Card6</div>
            </div>
          </div>
          <div className="extras-container extras">
            <div className="extras-title">Extras Title</div>
            <div className="mainextras-container">
              <div className="extra">Announcements</div>
              <div className="extra">Trending</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default App;
