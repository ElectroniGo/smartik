// See the Electron documentation for details on how to use preload scripts:
// https://www.electronjs.org/docs/latest/tutorial/process-model#preload-scripts

const { contextBridge, ipcRenderer } = require('electron');

// Expose protected methods that allow the renderer process to use
// the ipcRenderer without exposing the entire object
contextBridge.exposeInMainWorld('electronAPI', {
  // Window controls
  minimizeWindow: () => ipcRenderer.invoke('window:minimize'),
  maximizeWindow: () => ipcRenderer.invoke('window:maximize'),
  closeWindow: () => ipcRenderer.invoke('window:close'),
  
  // File operations
  openFile: () => ipcRenderer.invoke('file:open'),
  saveFile: (data) => ipcRenderer.invoke('file:save', data),
  
  // App information
  getAppVersion: () => ipcRenderer.invoke('app:version'),
  
  // Theme management
  setTheme: (theme) => ipcRenderer.invoke('theme:set', theme),
  getTheme: () => ipcRenderer.invoke('theme:get'),
  
  // Event listeners for renderer process
  onThemeChanged: (callback) => {
    ipcRenderer.on('theme:changed', callback);
    // Return a cleanup function
    return () => ipcRenderer.removeListener('theme:changed', callback);
  },
  
  // Notification system
  showNotification: (title, body) => ipcRenderer.invoke('notification:show', { title, body }),
  
  // System information
  getSystemInfo: () => ipcRenderer.invoke('system:info')
});

// Optional: Log when preload script is loaded (for debugging)
console.log('Preload script loaded successfully');