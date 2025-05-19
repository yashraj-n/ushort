document.addEventListener('DOMContentLoaded', function() {
  // Elements
  const longUrlInput = document.getElementById('long-url');
  const shortenBtn = document.getElementById('shorten-btn');
  const resultContainer = document.querySelector('.result-container');
  const shortUrlDisplay = document.getElementById('short-url');
  const copyBtn = document.getElementById('copy-btn');
  const qrBtn = document.getElementById('qr-btn');
  const qrContainer = document.querySelector('.qr-container');
  const qrCodeElement = document.getElementById('qrcode');
  const downloadQrBtn = document.getElementById('download-qr');
  const loadingSpinner = document.querySelector('.loading-spinner');
  const notification = document.getElementById('notification');
  const mobileMenuBtn = document.querySelector('.mobile-menu-btn');
  const mobileNav = document.querySelector('.mobile-nav');
  const mobileNavClose = document.querySelector('.mobile-nav-close');
  
  // Base URL for short links
  const baseUrl = 'https://linksnip.io/';
  
  // Random string generator (simulating a backend)
  function generateRandomString(length = 7) {
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let result = '';
    for (let i = 0; i < length; i++) {
      result += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return result;
  }
  
  // URL validation
  function isValidUrl(url) {
    try {
      new URL(url);
      return true;
    } catch (e) {
      return false;
    }
  }
  
  // Show notification
  function showNotification(message, type = 'success') {
    notification.textContent = message;
    notification.className = `notification ${type} show`;
    
    setTimeout(() => {
      notification.classList.remove('show');
    }, 3000);
  }
  
  // Create QR Code using QR Server API
  function generateQRCode(url) {
    // Clear previous QR code
    qrCodeElement.innerHTML = '';
    
    // Create an image element for the QR code
    const qrImage = document.createElement('img');
    qrImage.src = `https://api.qrserver.com/v1/create-qr-code/?size=160x160&data=${encodeURIComponent(url)}`;
    qrImage.alt = 'QR Code';
    qrImage.style.width = '160px';
    qrImage.style.height = '160px';
    
    qrCodeElement.appendChild(qrImage);
    
    // For downloading the QR code
    downloadQrBtn.onclick = () => {
      const link = document.createElement('a');
      link.href = qrImage.src;
      link.download = 'qrcode.png';
      link.click();
    };
  }
  
  // Shorten URL event handler
  shortenBtn.addEventListener('click', function() {
    const longUrl = longUrlInput.value.trim();
    
    if (!longUrl) {
      showNotification('Please enter a URL', 'error');
      return;
    }
    
    if (!isValidUrl(longUrl)) {
      showNotification('Please enter a valid URL', 'error');
      return;
    }
    
    // Show loading state
    shortenBtn.querySelector('span').style.display = 'none';
    loadingSpinner.style.display = 'block';
    
    // Simulate API call with setTimeout
    setTimeout(() => {
      const shortCode = generateRandomString();
      const shortUrl = baseUrl + shortCode;
      
      // Update the UI
      shortUrlDisplay.textContent = shortUrl;
      shortUrlDisplay.href = longUrl;
      resultContainer.style.display = 'block';
      
      // Hide loading state
      shortenBtn.querySelector('span').style.display = 'inline';
      loadingSpinner.style.display = 'none';
      
      // Generate QR code for the short URL
      generateQRCode(shortUrl);
      
      // Clear input
      longUrlInput.value = '';
      
      // Hide QR container initially
      qrContainer.style.display = 'none';
      
      showNotification('URL shortened successfully!');
    }, 1000);
  });
  
  // Copy button event handler
  copyBtn.addEventListener('click', function() {
    const textToCopy = shortUrlDisplay.textContent;
    
    navigator.clipboard.writeText(textToCopy)
      .then(() => {
        showNotification('Copied to clipboard!');
      })
      .catch(err => {
        showNotification('Failed to copy text', 'error');
      });
  });
  
  // QR code button event handler
  qrBtn.addEventListener('click', function() {
    qrContainer.style.display = qrContainer.style.display === 'none' ? 'flex' : 'none';
  });
  
  // Mobile menu handlers
  mobileMenuBtn.addEventListener('click', function() {
    mobileNav.classList.add('active');
  });
  
  mobileNavClose.addEventListener('click', function() {
    mobileNav.classList.remove('active');
  });
  
  // Enter key to submit
  longUrlInput.addEventListener('keypress', function(e) {
    if (e.key === 'Enter') {
      shortenBtn.click();
    }
  });
  
  // Scroll animation
  const fadeElements = document.querySelectorAll('.fade-in');
  
  const fadeInObserver = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('appear');
      }
    });
  }, {
    threshold: 0.1
  });
  
  fadeElements.forEach(element => {
    fadeInObserver.observe(element);
  });
  
  // Smooth scrolling for anchor links
  document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function(e) {
      if (this.getAttribute('href') === '#') return;
      
      e.preventDefault();
      
      const targetId = this.getAttribute('href');
      const targetElement = document.querySelector(targetId);
      
      if (targetElement) {
        targetElement.scrollIntoView({
          behavior: 'smooth'
        });
        
        // Close mobile menu if open
        mobileNav.classList.remove('active');
      }
    });
  });
}); 