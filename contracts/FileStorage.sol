pragma solidity >=0.4.0 <0.6.0;

contract FileStorage {
    //address of owner
    address public owner;
    
    //structure of file
    struct File {
        string name;
        string hash;
    }
    
    event added();
    //mapping of file
    mapping (address=>File) internal filePerUser;
    
    //
    constructor () public {
        owner = msg.sender;
        
    }
    
    function addFile(string memory hash,string memory name) public  {
      filePerUser[msg.sender].name = name;
      filePerUser[msg.sender].hash = hash;
     //return (filePerUser[msg.sender].hash);
    }
    
    // getHash of user
    function getHash() public  view returns(string memory){
        return filePerUser[msg.sender].hash;
    }
}