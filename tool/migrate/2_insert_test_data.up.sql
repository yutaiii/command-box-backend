INSERT INTO
command_box.commands(
    title,
    text
)
VALUES
(
    '安全にforce pushする',
    'git push --force-with-lease [remote-name] [branch-name]'
),
(
    'ファイルを部分的にaddする',
    'git add -p'
),
(
    'Docker Composeで作ったコンテナ、イメージ、ボリューム、ネットワークを一括完全消去する',
    'docker-compose down --rmi all --volumes --remove-orphans'
),
(
    'SLを走らせる',
    'sl'
),
(
    'hogeaaaaaaaaaaaaaaaaaaaa',
    'fugaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa'
);