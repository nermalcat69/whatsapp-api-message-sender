package handler

import (
	"html/template"
	"net/http"
)

var pageTmpl = template.Must(template.New("page").Parse(pageHTML))

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	pageTmpl.Execute(w, nil)
}

const pageHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>WhatsApp Sender</title>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      background: #f7f8fa; color: #111; min-height: 100vh;
    }
    header {
      background: #fff; border-bottom: 1px solid #e5e7eb;
      padding: 18px 32px; display: flex; align-items: center; gap: 10px;
    }
    header svg { color: #25D366; }
    header h1 { font-size: 16px; font-weight: 600; color: #111; letter-spacing: -0.3px; }
    .container { max-width: 1100px; margin: 0 auto; padding: 32px 24px; }
    .composer {
      background: #fff; border: 1px solid #e5e7eb;
      border-radius: 12px; padding: 24px; margin-bottom: 24px;
    }
    .composer-grid { display: grid; grid-template-columns: 1fr auto; gap: 16px; align-items: start; }
    label {
      display: block; font-size: 12px; font-weight: 500; color: #6b7280;
      text-transform: uppercase; letter-spacing: 0.5px; margin-bottom: 8px;
    }
    textarea {
      width: 100%; height: 140px; padding: 12px 14px;
      border: 1px solid #e5e7eb; border-radius: 8px; font-size: 14px;
      line-height: 1.6; color: #111; background: #fafafa;
      resize: vertical; outline: none; transition: border-color 0.15s;
    }
    textarea:focus { border-color: #25D366; background: #fff; }
    textarea::placeholder { color: #9ca3af; }
    .upload-zone {
      width: 200px; height: 140px; border: 1.5px dashed #d1d5db;
      border-radius: 8px; display: flex; flex-direction: column;
      align-items: center; justify-content: center; gap: 8px;
      cursor: pointer; transition: border-color 0.15s, background 0.15s;
      background: #fafafa; position: relative;
    }
    .upload-zone:hover { border-color: #25D366; background: #f0fdf4; }
    .upload-zone.has-file { border-color: #25D366; background: #f0fdf4; }
    .upload-zone input[type="file"] {
      position: absolute; inset: 0; opacity: 0;
      cursor: pointer; width: 100%; height: 100%;
    }
    .upload-zone .upload-icon { color: #9ca3af; }
    .upload-zone.has-file .upload-icon { color: #25D366; }
    .upload-zone span { font-size: 12px; color: #6b7280; text-align: center; padding: 0 12px; }
    .upload-zone.has-file span { color: #15803d; font-weight: 500; }
    .hint { font-size: 12px; color: #9ca3af; margin-top: 8px; }
    .hint code {
      background: #f3f4f6; padding: 1px 5px; border-radius: 4px;
      font-family: "SF Mono", monospace; font-size: 11px; color: #374151;
    }
    .section-header { display: flex; align-items: center; justify-content: space-between; margin-bottom: 16px; }
    .section-header h2 { font-size: 14px; font-weight: 600; color: #374151; }
    .badge {
      font-size: 11px; font-weight: 500; background: #f3f4f6;
      color: #6b7280; padding: 2px 8px; border-radius: 20px;
    }
    .cards-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(300px, 1fr)); gap: 14px; }
    .card {
      background: #fff; border: 1px solid #e5e7eb; border-radius: 10px;
      padding: 18px; display: flex; flex-direction: column; gap: 12px;
      transition: box-shadow 0.15s;
    }
    .card:hover { box-shadow: 0 2px 8px rgba(0,0,0,0.06); }
    .card-top { display: flex; align-items: center; gap: 12px; }
    .avatar {
      width: 38px; height: 38px; border-radius: 50%;
      background: #dcfce7; color: #15803d; font-size: 14px; font-weight: 600;
      display: flex; align-items: center; justify-content: center;
      flex-shrink: 0; text-transform: uppercase;
    }
    .contact-info { flex: 1; min-width: 0; }
    .contact-name { font-size: 14px; font-weight: 600; color: #111; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
    .contact-phone { font-size: 12px; color: #9ca3af; margin-top: 1px; }
    .message-preview {
      background: #f9fafb; border: 1px solid #f3f4f6; border-radius: 8px;
      padding: 10px 12px; font-size: 13px; color: #374151; line-height: 1.55;
      white-space: pre-wrap; word-break: break-word; max-height: 90px;
      overflow: hidden; position: relative;
    }
    .message-preview::after {
      content: ''; position: absolute; bottom: 0; left: 0; right: 0; height: 28px;
      background: linear-gradient(transparent, #f9fafb);
      border-radius: 0 0 8px 8px; pointer-events: none;
    }
    .send-btn {
      display: inline-flex; align-items: center; justify-content: center;
      gap: 7px; padding: 9px 16px; background: #25D366; color: #fff;
      border: none; border-radius: 8px; font-size: 13px; font-weight: 500;
      cursor: pointer; text-decoration: none;
      transition: background 0.15s, transform 0.1s; width: 100%;
    }
    .send-btn:hover { background: #1ebe5d; }
    .send-btn:active { transform: scale(0.98); }
    .empty-state { text-align: center; padding: 60px 0; color: #9ca3af; }
    .empty-state svg { margin-bottom: 12px; opacity: 0.4; }
    .empty-state p { font-size: 14px; }
    .send-all-bar {
      display: none; align-items: center; justify-content: space-between;
      background: #fff; border: 1px solid #e5e7eb; border-radius: 10px;
      padding: 14px 20px; margin-bottom: 20px;
    }
    .send-all-bar.visible { display: flex; }
    .send-all-bar p { font-size: 13px; color: #374151; }
    .send-all-bar p strong { font-weight: 600; }
    .send-all-btn {
      display: inline-flex; align-items: center; gap: 7px;
      padding: 8px 18px; background: #25D366; color: #fff; border: none;
      border-radius: 8px; font-size: 13px; font-weight: 500;
      cursor: pointer; transition: background 0.15s;
    }
    .send-all-btn:hover { background: #1ebe5d; }
  </style>
</head>
<body>

<header>
  <svg width="22" height="22" fill="currentColor" viewBox="0 0 24 24">
    <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/>
  </svg>
  <h1>WhatsApp Sender</h1>
</header>

<div class="container">
  <div class="composer">
    <div class="composer-grid">
      <div>
        <label>Message Template</label>
        <textarea id="messageTemplate" placeholder="Hey {contactName}, just wanted to share something with you..."></textarea>
        <p class="hint">Use <code>{contactName}</code> to personalize each message.</p>
      </div>
      <div>
        <label>Contacts (VCF)</label>
        <div class="upload-zone" id="uploadZone">
          <input type="file" id="vcfInput" accept=".vcf,text/vcard" />
          <svg class="upload-icon" width="24" height="24" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5m-13.5-9L12 3m0 0l4.5 4.5M12 3v13.5"/>
          </svg>
          <span id="uploadLabel">Upload .vcf file</span>
        </div>
      </div>
    </div>
  </div>

  <div class="send-all-bar" id="sendAllBar">
    <p>Ready to send to <strong id="contactCount">0</strong> contacts</p>
    <button class="send-all-btn" id="sendAllBtn">
      <svg width="15" height="15" fill="currentColor" viewBox="0 0 24 24">
        <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/>
      </svg>
      Send All
    </button>
  </div>

  <div class="section-header" id="cardsHeader" style="display:none">
    <h2>Contacts</h2>
    <span class="badge" id="cardsBadge">0 contacts</span>
  </div>

  <div class="cards-grid" id="cardsGrid"></div>

  <div class="empty-state" id="emptyState">
    <svg width="48" height="48" fill="none" stroke="currentColor" stroke-width="1" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" d="M18 18.72a9.094 9.094 0 003.741-.479 3 3 0 00-4.682-2.72m.94 3.198l.001.031c0 .225-.012.447-.037.666A11.944 11.944 0 0112 21c-2.17 0-4.207-.576-5.963-1.584A6.062 6.062 0 016 18.719m12 0a5.971 5.971 0 00-.941-3.197m0 0A5.995 5.995 0 0012 12.75a5.995 5.995 0 00-5.058 2.772m0 0a3 3 0 00-4.681 2.72 8.986 8.986 0 003.74.477m.94-3.197a5.971 5.971 0 00-.94 3.197M15 6.75a3 3 0 11-6 0 3 3 0 016 0zm6 3a2.25 2.25 0 11-4.5 0 2.25 2.25 0 014.5 0zm-13.5 0a2.25 2.25 0 11-4.5 0 2.25 2.25 0 014.5 0z"/>
    </svg>
    <p>Upload a VCF file to load your contacts</p>
  </div>
</div>

<script>
  var contacts = [];

  var messageTemplate = document.getElementById('messageTemplate');
  var vcfInput        = document.getElementById('vcfInput');
  var uploadZone      = document.getElementById('uploadZone');
  var uploadLabel     = document.getElementById('uploadLabel');
  var cardsGrid       = document.getElementById('cardsGrid');
  var emptyState      = document.getElementById('emptyState');
  var cardsHeader     = document.getElementById('cardsHeader');
  var cardsBadge      = document.getElementById('cardsBadge');
  var sendAllBar      = document.getElementById('sendAllBar');
  var contactCount    = document.getElementById('contactCount');
  var sendAllBtn      = document.getElementById('sendAllBtn');

  function formatMessage(tmpl, name) {
    return tmpl.replace(/\{contactName\}/g, name);
  }

  function initials(name) {
    return name.split(' ').map(function(w) { return w[0]; }).slice(0, 2).join('');
  }

  function buildWhatsAppURL(phone, message) {
    var cleaned = phone.replace(/\D/g, '');
    return 'whatsapp://send?phone=' + encodeURIComponent(cleaned) + '&text=' + encodeURIComponent(message);
  }

  function escapeHTML(str) {
    return str
      .replace(/&/g, '&amp;')
      .replace(/</g, '&lt;')
      .replace(/>/g, '&gt;')
      .replace(/"/g, '&quot;');
  }

  function buildCardHTML(c, idx) {
    var msg = formatMessage(messageTemplate.value, c.name);
    var url = buildWhatsAppURL(c.phone, msg);
    var preview = msg ? escapeHTML(msg) : '<span style="color:#9ca3af;font-style:italic">No message yet</span>';
    var waIcon = '<svg width="14" height="14" fill="currentColor" viewBox="0 0 24 24"><path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/></svg>';

    return '<div class="card" data-idx="' + idx + '">'
      + '<div class="card-top">'
      +   '<div class="avatar">' + initials(c.name) + '</div>'
      +   '<div class="contact-info">'
      +     '<div class="contact-name">' + escapeHTML(c.name) + '</div>'
      +     '<div class="contact-phone">' + escapeHTML(c.phone) + '</div>'
      +   '</div>'
      + '</div>'
      + '<div class="message-preview">' + preview + '</div>'
      + '<a class="send-btn" href="' + url + '" id="send-' + idx + '">' + waIcon + ' Open in WhatsApp</a>'
      + '</div>';
  }

  function renderCards() {
    cardsGrid.innerHTML = '';

    if (contacts.length === 0) {
      emptyState.style.display = 'block';
      cardsHeader.style.display = 'none';
      sendAllBar.classList.remove('visible');
      return;
    }

    emptyState.style.display = 'none';
    cardsHeader.style.display = 'flex';
    sendAllBar.classList.add('visible');
    cardsBadge.textContent = contacts.length + ' contact' + (contacts.length !== 1 ? 's' : '');
    contactCount.textContent = contacts.length;

    cardsGrid.innerHTML = contacts.map(function(c, idx) {
      return buildCardHTML(c, idx);
    }).join('');
  }

  messageTemplate.addEventListener('input', function() {
    if (contacts.length > 0) renderCards();
  });

  vcfInput.addEventListener('change', function() {
    var file = vcfInput.files[0];
    if (!file) return;
    uploadLabel.textContent = 'Uploading...';

    var formData = new FormData();
    formData.append('vcf', file);

    fetch('/parse-vcf', { method: 'POST', body: formData })
      .then(function(res) {
        if (!res.ok) throw new Error('Upload failed');
        return res.json();
      })
      .then(function(data) {
        contacts = data || [];
        uploadZone.classList.add('has-file');
        uploadLabel.textContent = file.name + ' (' + contacts.length + ')';
        renderCards();
      })
      .catch(function(err) {
        uploadLabel.textContent = 'Error — try again';
        uploadZone.classList.remove('has-file');
        console.error(err);
      });
  });

  sendAllBtn.addEventListener('click', function() {
    contacts.forEach(function(c, idx) {
      setTimeout(function() {
        var msg = formatMessage(messageTemplate.value, c.name);
        window.location.href = buildWhatsAppURL(c.phone, msg);
      }, idx * 1200);
    });
  });

  renderCards();
</script>
</body>
</html>`
