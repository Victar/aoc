package adventofcode.year2020;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;
import lombok.ToString;

public class Day19 extends BaseTest {

	@Test public void singleCheck() {
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day19/input.txt");
		final ArrayList<Rule> rulesList = new ArrayList<>();
		final ArrayList<String> messageList = new ArrayList<>();
		fillData(data, rulesList, messageList);
		rulesList.sort(Comparator.comparing(a -> a.getNum()));
		final Rule rule0 = rulesList.stream().filter(l -> l.getNum() == 0).findFirst().orElse(null);
		final String regExp0 = rule0.toRegExpString();
		//		System.out.println(regExp0);
		int count = 0;
		for (final String message : messageList) {
			if (message.matches(regExp0)) {
				count++;
			}
		}
		System.out.println(count);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day19/input_replace.txt");
		final ArrayList<Rule> rulesList = new ArrayList<>();
		final ArrayList<String> messageList = new ArrayList<>();
		fillData(data, rulesList, messageList);
		rulesList.sort(Comparator.comparing(a -> a.getNum()));
		final Rule rule0 = rulesList.stream().filter(l -> l.getNum() == 0).findFirst().orElse(null);
		final String regExp0 = rule0.toRegExpString();
		//		System.out.println(regExp0);
		int count = 0;
		for (final String message : messageList) {
			if (message.matches(regExp0)) {
				count++;
			}
		}
		System.out.println(count);
	}

	public void fillData(final ArrayList<String> data, final ArrayList<Rule> rulesList, final ArrayList<String> messageList) {
		for (final String input : data) {
			if (input.contains(":")) {
				//parse rule
				final String[] inputArray = StringUtils.split(input, ":\\|");
				final int ruleNumber = Integer.parseInt(inputArray[0]);
				Rule currentRule = rulesList.stream().filter(l -> l.getNum() == ruleNumber).findFirst().orElse(null);
				if (currentRule == null) {
					currentRule = new Rule(ruleNumber);
					rulesList.add(currentRule);
				}
				final String input2 = inputArray[1];
				if (input2.contains("a")) {
					currentRule.setC("a");
				} else if (input2.contains("b")) {
					currentRule.setC("b");
				} else {
					final String[] childLeft = inputArray[1].split(" ");
					for (int i = 0; i < childLeft.length; i++) {
						if (StringUtils.isNotEmpty(childLeft[i])) {
							final int childRuleNumber = Integer.parseInt(childLeft[i].trim());
							Rule childRule = rulesList.stream().filter(l -> l.getNum() == childRuleNumber).findFirst().orElse(null);
							if (childRule == null) {
								childRule = new Rule(childRuleNumber);
								rulesList.add(childRule);
							}
							currentRule.addLeft(childRule);
						}
					}
				}
				if (inputArray.length == 3) {
					final String[] childRight = inputArray[2].split(" ");
					for (int i = 0; i < childRight.length; i++) {
						if (StringUtils.isNotEmpty(childRight[i])) {

							final int childRuleNumber = Integer.parseInt(childRight[i].trim());
							Rule childRule = rulesList.stream().filter(l -> l.getNum() == childRuleNumber).findFirst().orElse(null);
							if (childRule == null) {
								childRule = new Rule(childRuleNumber);
								rulesList.add(childRule);
							}
							currentRule.addRight(childRule);
						}
					}
				}
			} else {
				if (StringUtils.isNotEmpty(input)) {
					messageList.add(input);
				}
			}
		}
	}

	@Data @ToString class Rule {

		int num;
		List<Rule> left = new ArrayList<>();
		List<Rule> right = new ArrayList<>();
		String c;

		int goldVisitedTime;

		public Rule(final int num) {
			this.num = num;
		}

		public String toRegExpString() {
			if (StringUtils.isNotEmpty(this.c)) {
				return this.c;
			}
			final StringBuilder sb = new StringBuilder("(");
			for (final Rule rule : this.left) {
				if (this.num != rule.num || this.goldVisitedTime < 10) {
					if (this.num == rule.num) {
						this.goldVisitedTime++;
					}
					sb.append("(" + rule.toRegExpString() + ")");
				}
			}
			sb.append(")");
			if (this.right.size() > 0) {
				sb.append("|(");
				for (final Rule rule : this.right) {
					if (this.num != rule.num || this.goldVisitedTime < 10) {
						if (this.num == rule.num) {
							this.goldVisitedTime++;
						}
						sb.append("(" + rule.toRegExpString() + ")");
					}
				}
				sb.append(")");
			}
			return sb.toString();
		}

		public void addLeft(final Rule rule) {
			this.left.add(rule);
		}

		public void addRight(final Rule rule) {
			this.right.add(rule);
		}

		public void setC(final String c) {
			this.c = c;
		}

		@Override public String toString() {
			return "Rule{" + "num=" + this.num + ", left=" + this.left.size() + ", right=" + this.right.size() + ", c='" + this.c + '\''
					+ '}';
		}
	}
}

