-- +goose Up

--
-- Дамп данных таблицы `product`
--

INSERT INTO `product` (`id`, `Name`, `Image_name`, `Height`, `Preparation`) VALUES
(0, 'фикус', 'fikus.jpg', 15, 1),
(1, 'куст', 'kust.jpg', 10, 1),
(2, 'дерево', 'derevo.jpg', 25, 3);

--
-- Дамп данных таблицы `product_tag`
--

INSERT INTO `product_tag` (`product_id`, `tag_id`) VALUES
(1, 2),
(0, 1),
(0, 0),
(2, 1),
(2, 0),
(2, 1);

--
-- Дамп данных таблицы `tag`
--

INSERT INTO `tag` (`id`, `Name`) VALUES
(0, 'острый'),
(1, 'необычный'),
(2, 'обычный');

-- +goose Down
