const { exec } = require('child_process');

const killNodeProcesses = () => {
    // Get PIDs of Node.js processes
    exec('tasklist | findstr node.exe', (error, stdout) => {
        if (error) {
            console.error(`Error: ${error.message}`);
            return;
        }
        
        const PIDs = stdout.split(/\s+/).filter(pid => !isNaN(parseInt(pid)));
        
        // Kill Node.js processes using their PIDs
        PIDs.forEach(pid => {
            exec(`taskkill /PID ${pid} /F`, (error) => {
                if (error) {
                    console.error(`Error killing process with PID ${pid}: ${error.message}`);
                } else {
                    console.log(`Process with PID ${pid} killed successfully`);
                }
            });
        });
    });
};

killNodeProcesses();