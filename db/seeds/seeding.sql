-- ============================================================
-- SEEDER: shortlink_db
-- 30 users + 30 links
-- Password hash: argon2id("pass1234") — same for all seed users
-- Params: m=65536 (64MiB), t=2, p=1, keyLen=32, saltLen=16 (sesuai HashConfig OWASP)
-- ============================================================

-- ============================================================
-- SEED: users (30 rows)
-- ============================================================
INSERT INTO users (email, password_hash, full_name, avatar_url, created_at) VALUES
('ahmad.fauzi@gmail.com',      '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Ahmad Fauzi',        'https://i.pravatar.cc/150?img=1',  '2025-01-05 08:12:00'),
('budi.santoso@yahoo.com',     '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Budi Santoso',       'https://i.pravatar.cc/150?img=2',  '2025-01-10 09:30:00'),
('citra.dewi@outlook.com',     '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Citra Dewi',         'https://i.pravatar.cc/150?img=3',  '2025-01-15 10:45:00'),
('deni.rahman@gmail.com',      '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Deni Rahman',        'https://i.pravatar.cc/150?img=4',  '2025-01-20 11:00:00'),
('eka.putri@gmail.com',        '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Eka Putri',          'https://i.pravatar.cc/150?img=5',  '2025-02-01 07:20:00'),
('fajar.nugroho@hotmail.com',  '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Fajar Nugroho',      'https://i.pravatar.cc/150?img=6',  '2025-02-05 13:15:00'),
('gita.lestari@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Gita Lestari',       'https://i.pravatar.cc/150?img=7',  '2025-02-10 14:00:00'),
('hendra.wijaya@gmail.com',    '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Hendra Wijaya',      'https://i.pravatar.cc/150?img=8',  '2025-02-15 08:45:00'),
('indah.sari@gmail.com',       '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Indah Sari',         'https://i.pravatar.cc/150?img=9',  '2025-02-20 16:30:00'),
('joko.susilo@yahoo.com',      '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Joko Susilo',        'https://i.pravatar.cc/150?img=10', '2025-03-01 09:00:00'),
('kartika.sari@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Kartika Sari',       'https://i.pravatar.cc/150?img=11', '2025-03-05 10:20:00'),
('lukman.hakim@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Lukman Hakim',       'https://i.pravatar.cc/150?img=12', '2025-03-10 11:45:00'),
('maya.anggraini@outlook.com', '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Maya Anggraini',     'https://i.pravatar.cc/150?img=13', '2025-03-15 13:00:00'),
('nanda.pratama@gmail.com',    '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Nanda Pratama',      'https://i.pravatar.cc/150?img=14', '2025-03-20 14:30:00'),
('okta.wulandari@gmail.com',   '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Okta Wulandari',     'https://i.pravatar.cc/150?img=15', '2025-04-01 08:00:00'),
('pandu.kusuma@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Pandu Kusuma',       'https://i.pravatar.cc/150?img=16', '2025-04-05 09:15:00'),
('qori.amelia@yahoo.com',      '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Qori Amelia',        'https://i.pravatar.cc/150?img=17', '2025-04-10 10:30:00'),
('reza.firmansyah@gmail.com',  '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Reza Firmansyah',    'https://i.pravatar.cc/150?img=18', '2025-04-15 11:45:00'),
('siska.permata@gmail.com',    '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Siska Permata',      'https://i.pravatar.cc/150?img=19', '2025-04-20 13:00:00'),
('taufik.hidayat@gmail.com',   '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Taufik Hidayat',     'https://i.pravatar.cc/150?img=20', '2025-05-01 07:30:00'),
('umar.saputra@hotmail.com',   '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Umar Saputra',       'https://i.pravatar.cc/150?img=21', '2025-05-05 08:45:00'),
('vera.novita@gmail.com',      '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Vera Novita',        'https://i.pravatar.cc/150?img=22', '2025-05-10 10:00:00'),
('wahyu.setiawan@gmail.com',   '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Wahyu Setiawan',     'https://i.pravatar.cc/150?img=23', '2025-05-15 11:15:00'),
('xena.clara@outlook.com',     '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Xena Clara',         'https://i.pravatar.cc/150?img=24', '2025-05-20 12:30:00'),
('yudi.prasetyo@gmail.com',    '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Yudi Prasetyo',      'https://i.pravatar.cc/150?img=25', '2025-06-01 08:00:00'),
('zara.aulia@gmail.com',       '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Zara Aulia',         'https://i.pravatar.cc/150?img=26', '2025-06-05 09:20:00'),
('aldi.mukhlis@yahoo.com',     '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Aldi Mukhlis',       'https://i.pravatar.cc/150?img=27', '2025-06-10 10:40:00'),
('bella.oktavia@gmail.com',    '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Bella Oktavia',      'https://i.pravatar.cc/150?img=28', '2025-06-15 11:55:00'),
('candra.purnama@gmail.com',   '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Candra Purnama',     'https://i.pravatar.cc/150?img=29', '2025-07-01 07:10:00'),
('dinda.rahayu@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$l+BlAj3I08uzOAEzu6/TAA$EG7AkyYgsAATdA5x9etUB3GLoObOfbya49IESaMo9D0', 'Dinda Rahayu',       'https://i.pravatar.cc/150?img=30', '2025-07-05 08:25:00');

-- Reset sequence after manual inserts (if needed)
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));


-- ============================================================
-- SEED: links (30 rows)
-- user_id=1 (Ahmad Fauzi) memiliki 12 link
-- user_id=13–30 masing-masing 1 link (18 link)
-- 5 link di-soft-delete (deleted_at IS NOT NULL)
-- ============================================================
INSERT INTO links (user_id, original_url, slug, click_count, created_at, deleted_at) VALUES
-- Ahmad Fauzi (user_id=1) — 12 links ---
(1,  'https://www.tokopedia.com/promo/harbolnas-2025',           'tokped-hrb',    142, '2025-01-06 09:00:00', NULL),
(1,  'https://www.shopee.co.id/flash-sale/januari',              'shp-flash',     389, '2025-01-11 10:00:00', NULL),
(1,  'https://docs.google.com/spreadsheets/d/abcdef123456/edit', 'gdoc-sheet',     57, '2025-01-16 11:00:00', NULL),
(1,  'https://www.youtube.com/watch?v=dQw4w9WgXcQ',             'yt-rkroll',    9021, '2025-01-21 12:00:00', NULL),
(1,  'https://github.com/laravel/laravel',                       'gh-laravel',    204, '2025-02-02 08:00:00', NULL),
(1,  'https://www.figma.com/file/xyz789/design-system',          'fig-ds',         88, '2025-02-06 14:00:00', NULL),
(1,  'https://medium.com/@gitadev/tips-react-hooks-2025',        'med-react',     312, '2025-02-11 15:00:00', NULL),
(1,  'https://www.canva.com/design/DAF000/view',                 'cnv-poster',     45, '2025-02-16 09:30:00', NULL),
(1,  'https://trello.com/b/abc123/project-sprint-5',             'trl-sprint5',    76, '2025-02-21 17:00:00', NULL),
(1,  'https://notion.so/team/roadmap-q1-2025',                   'ntn-roadmap',   133, '2025-03-02 10:00:00', NULL),
(1,  'https://www.detik.com/teknologi/berita-terbaru',           'dtk-tech',      567, '2025-03-06 11:00:00', NULL),
(1,  'https://stackoverflow.com/questions/tagged/postgresql',    'so-postgres',   220, '2025-03-11 12:00:00', NULL),
-- Other users — 1 link each ---
(13, 'https://www.npmjs.com/package/express',                    'npm-express',   189, '2025-03-16 13:00:00', NULL),
(14, 'https://laravel.com/docs/11.x/eloquent',                   'lrv-eloquent',   97, '2025-03-21 14:00:00', NULL),
(15, 'https://tailwindcss.com/docs/installation',                'tw-install',    443, '2025-04-02 09:00:00', NULL),
(16, 'https://www.linkedin.com/in/pandukusuma',                  'li-pandu',       61, '2025-04-06 10:00:00', NULL),
(17, 'https://drive.google.com/drive/folders/proposal-ta-2025', 'gdrv-ta',        29, '2025-04-11 11:00:00', NULL),
(18, 'https://www.instagram.com/p/C4kLmNoPqRs/',                'ig-post-rz',    874, '2025-04-16 12:00:00', NULL),
(19, 'https://web.whatsapp.com',                                 'wa-web',       1502, '2025-04-21 13:00:00', NULL),
(20, 'https://meet.google.com/abc-defg-hij',                     'gmeet-daily',   308, '2025-05-02 08:00:00', NULL),
-- Soft-deleted links (deleted_at IS NOT NULL) ---
(21, 'https://www.bukalapak.com/deals/old-promo',                'bl-old-deal',    14, '2025-05-06 09:00:00', '2025-05-10 18:00:00'),
(22, 'https://www.gojek.com/promo/expired-voucher',              'gojek-exp',       3, '2025-05-11 10:00:00', '2025-05-15 20:00:00'),
(23, 'https://forms.gle/old-survey-form-xyz',                    'gform-old',       8, '2025-05-16 11:00:00', '2025-05-20 17:00:00'),
(24, 'https://bit.ly/redirect-removed',                          'blly-gone',       0, '2025-05-21 12:00:00', '2025-05-25 09:00:00'),
(25, 'https://www.eventbrite.com/e/past-event-123456',           'evbrt-past',     22, '2025-06-02 08:00:00', '2025-06-05 16:00:00'),
-- Active links continued ---
(26, 'https://chat.openai.com/share/xyz-session',                'oai-share',     658, '2025-06-06 10:00:00', NULL),
(27, 'https://www.udemy.com/course/nodejs-complete-guide/',      'udy-nodejs',    124, '2025-06-11 11:00:00', NULL),
(28, 'https://www.coursera.org/learn/machine-learning',          'crs-ml',        271, '2025-06-16 12:00:00', NULL),
(29, 'https://vercel.com/candra/my-nextjs-app/deployments',     'vcl-deploy',     48, '2025-07-02 09:00:00', NULL),
(30, 'https://railway.app/project/dinda-shortlink-backend',      'rail-backend',   91, '2025-07-06 10:00:00', NULL);

-- Reset sequence
SELECT setval('links_id_seq', (SELECT MAX(id) FROM links));